---
title: Autenticação com Cognito
sidebar_position: 1
slug: /autenticação2
---

## Introdução

Esta seção do documento apresenta uma descrição detalhada da implementação de um sistema de autenticação utilizando o serviço AWS Cognito, bem como sua relevância na proteção das rotas da aplicação com a interface e resultados da simulação de sensores para vigilância ambiental na cidade de São Paulo. Com foco na segurança e integridade dos dados, exploraremos o funcionamento dos códigos desenvolvidos, destacando a importância da autenticação com o AWS Cognito não apenas na proteção das rotas de login do usuário, mas também na preservação da integridade dos dados ambientais coletados e visualizados através de um dashboard. Ao longo deste documento, examinaremos como o uso do AWS Cognito fortalece a segurança do sistema, garantindo acesso apenas a usuários autenticados e autorizados, contribuindo assim para uma vigilância ambiental eficaz na metrópole paulista.

## O que é o Cognito?

O Cognito é uma plataforma da aws responsável por armazenar e validar dados de acesso de usuários. Com o cognito é possível cadastrar usuário e armazenar as suas informações. Além de gerar tokens do tipo OAuth, o cognito também consegue prover toda a validação de usuários.
Conseguimos armazenar alguns dados do usuário como: email, name, phone, birthdate, nickname, gender, website e muitos outros, também conseguimos colocar campos personalizados.
O cognito ainda nos permite trabalhar com "provedores federados", conhecidos como login social, como Google, Facebook e GitHub, o que não será abordado inicialmente na nossa aplicação, mas é possível de fazer com o cognito.

## Funcionamento dos Códigos

### `cognito.go`

O arquivo `cognito.go` encapsula a lógica de interação com o serviço AWS Cognito, fornecendo uma estrutura organizada para lidar com operações de autenticação de usuários. Este arquivo é essencial para a integração do sistema com o AWS Cognito e inclui as seguintes funções principais:

- **`SignUp`**:

Esta função permite registrar novos usuários no sistema. Ela recebe os dados do usuário, como nome, email e senha, e os envia para o AWS Cognito para criação da conta.

- **`ConfirmAccount`**:

Após o registro, os usuários precisam confirmar suas contas por meio de um código enviado por email. Esta função confirma a conta do usuário no AWS Cognito com base no código recebido.

- **`SignIn`**:

Responsável por autenticar usuários existentes. Recebe as credenciais do usuário (email e senha), valida-as junto ao AWS Cognito e retorna um token de acesso válido em caso de sucesso.

-  **`GetUserByToken`**:

Função de teste que permite obter informações do usuário com base em um token de acesso válido. Esta função é útil para recuperar dados do usuário autenticado, mas é fundamentalmente importante para validar a proteção de uma rota a partir do token gerado.

### `main.go`

O arquivo `main.go` atua como o ponto de entrada da aplicação e define as rotas HTTP para interação com o sistema de autenticação. Ele é responsável por direcionar as solicitações dos clientes para as funções adequadas do arquivo `cognito.go`. As principais rotas definidas neste arquivo incluem:

- **`POST /user`**:

Rota para registrar novos usuários. Recebe os dados do usuário via solicitação HTTP, chama a função `SignUp` do `cognito.go` para criar a conta e retorna uma resposta adequada ao cliente.

- **`POST /user/confirmation`**:

Rota para confirmar contas de usuários após o registro. Recebe o código de confirmação via solicitação HTTP, chama a função `ConfirmAccount` do `cognito.go` e retorna uma resposta adequada ao cliente.

- **`POST /user/login`**:

Rota para autenticar usuários existentes. Recebe as credenciais de login via solicitação HTTP, chama a função `SignIn` do `cognito.go` e retorna um token de acesso em caso de sucesso.

- **`GET /user`**:

Rota de teste para obter informações do usuário autenticado. Utiliza o token de acesso presente no cabeçalho da solicitação HTTP para chamar a função `GetUserByToken` do `cognito.go` e retornar os dados do usuário autenticado. Essa rota foi implementada apenas como maneira de testar o funcionamento da proteção das rotas com a autorização via token.

## Importância da Autenticação com o Cognito no nosso sistema

### Proteção das rotas de login do usuário

A utilização do AWS Cognito para autenticação proporciona uma camada robusta de segurança para as rotas de login do usuário. Esta segurança é assegurada através de três principais fatores, sendo eles a Segurança Avançada proporcionada pelo serviço da AWS, pois, uma vez que o AWS Cognito gerencia a autenticação de usuários de forma segura, ele implementa técnicas robustas de criptografia e autenticação de dois fatores e isso garante que as credenciais dos usuários sejam protegidas contra ameaças de segurança, como ataques de força bruta e roubo de identidade. Além disso, há a Facilidade de Integração com o serviço, uma vez que o SDK do AWS Cognito simplifica a integração da autenticação em aplicativos web e móveis e com uma API simples e eficiente, os desenvolvedores podem facilmente incorporar funcionalidades de autenticação em seus aplicativos, economizando tempo e esforço de desenvolvimento. Por fim, a Escalabilidade, já que o serviço Cognito é altamente escalável e pode lidar com milhões de usuários simultâneos sem comprometer a segurança ou o desempenho. Isso é essencial para garantir que a aplicação possa crescer e atender a demandas crescentes de usuários sem sacrificar a qualidade ou a segurança da autenticação.

### Proteção de rotas do backend de Sensores

Além de proteger as rotas de login do usuário na nossa aplicação, a autenticação com o AWS Cognito pode ser estendida para proteger as rotas do backend de simulação de sensores. Esta medida é crucial para garantir que apenas usuários autenticados e autorizados possam acessar e manipular os dados sensíveis coletados pelos sensores. Ao restringir o acesso às rotas do backend de sensores, o sistema pode prevenir potenciais tentativas de manipulação ou adulteração dos dados, garantindo assim a integridade e a confiabilidade das informações ambientais coletadas.

## Integração

Para a integração com o frontend em Next.js da nossa aplicação, foi utilizada a API nativa do navegador, Fetch. A integração permite o envio de requisições HTTP do front-end para o backend, possibilitando operações como cadastro de usuários e confirmação de contas.

### Passo a passo da integração

1. **Definição do Formulário de Cadastro de Usuário:**

O front-end Next.js inclui um formulário para cadastro de novos usuários, contendo campos para nome, e-mail, senha e confirmação de senha.

2. **Tratamento do Evento de Envio do Formulário:**

Quando o usuário preenche o formulário e clica em "Submit", é acionado o evento onFinish do formulário, que chama a função `handleSignUp`.

3. **Envio da Requisição de Cadastro para o Backend:**

A função `handleSignUp` é responsável por fazer uma requisição POST para o endpoint `/user` do backend, contendo os dados do novo usuário (nome, e-mail, senha). Assim, a requisição é realizada utilizando a função `fetch`, uma API nativa do navegador para fazer requisições HTTP. Para o envio das infirmações, o método utilizado é POST, e o corpo da requisição é um objeto JSON com os dados do usuário.

4. **Tratamento da Resposta do Backend:**

Se a requisição for bem-sucedida `(status 200 OK)`, o backend cria o usuário e retorna uma resposta com os dados do usuário recém-criado. Caso ocorra algum erro na requisição, como o e-mail já estar em uso, é lançada uma exceção que é capturada pelo bloco `catch`, onde é exibido um erro no console.

5. **Modal de Verificação de Conta:**

Após o cadastro do usuário ser realizado com sucesso, é exibido um modal de verificação de conta, solicitando que o usuário insira o código de verificação enviado para o seu e-mail.

6. **Envio da Requisição de Confirmação para o Backend:**

Ao clicar em "Enviar Código" no modal, é acionado o evento `onOk`, que chama a função `handleOk`. A função `handleOk` faz uma nova requisição POST para o endpoint `/user/confirmation` do backend, passando o e-mail do usuário e o código de verificação. Por fim, a requisição é feita utilizando a função `fetch`, com método POST e corpo da requisição em JSON.

7. **Tratamento da Resposta do Backend:**

Se a requisição de confirmação for bem-sucedida, o usuário é redirecionado para a página de login. Caso contrário, é exibida uma mensagem de erro no console.

## Conclusão

Em suma, a implementação da autenticação com o AWS Cognito representa um passo fundamental na construção de sistemas seguros e confiáveis. Ao garantir a autenticidade dos usuários e proteger as rotas de acesso, o Cognito desempenha um papel crucial na preservação da integridade dos dados e na prevenção de acessos não autorizados. Além disso, ao estender essa proteção para o backend de sensores, o sistema assegura que apenas usuários autenticados possam acessar e manipular informações sensíveis, promovendo assim a confiabilidade dos dados ambientais coletados. Além disso, a integração do front-end Next.js com o backend usando Fetch permite que o front-end envie requisições HTTP para o backend e trate as respostas recebidas, possibilitando operações de cadastro de usuários e confirmação de contas. Com isso, a vigilância ambiental na cidade de São Paulo é fortalecida, permitindo intervenções mais ágeis e eficazes para garantir a qualidade do ar e do ambiente, contribuindo para uma cidade mais saudável e sustentável para seus habitantes.

## Demo

<iframe width="560" height="315" src="https://www.youtube.com/embed/ksPyfoVVPQ8?si=VPRYxyKxNueYIY_v" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>