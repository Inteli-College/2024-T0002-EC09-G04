---
title: Autenticação do Sistema
sidebar_position: 1
slug: /autenticação
---

## Introdução

No contexto do desenvolvimento de uma plataforma para simulação de sensores IoT para ecovigilância em cidades inteligentes, a questão da autenticação e autorização desempenha um papel crucial. Para que haja um acesso mais seguro dos usuários à plataforma com dashboards da aplicação com dados dos sensores de ecovigilância na cidade de São Paulo, tecnologias como "JWT" e "OAuth 2.0" para autenticação e autorização, além de fornecer insights valiosos para tomadores de decisão e autoridades municipais, garantem que apenas usuários autorizados, como colaboradores e administradores, tenham acesso aos dados sensíveis e possam interagir com certas funcionalidades da aplicação.

## Garantia da Confidencialidade

A confidencialidade é um dos princípios fundamentais da segurança da informação e da triade CIA (Confidentiality, Integrity, Availability). Garantir a confidencialidade significa proteger as informações contra acesso não autorizado.

- Tokenização Segura: Tanto JWT quanto OAuth 2.0 oferecem métodos seguros de tokenização para autenticação e autorização. Os tokens JWT são assinados digitalmente pelo servidor, enquanto os tokens OAuth 2.0 são emitidos pelo provedor de serviços de autorização após a autenticação bem-sucedida do usuário.

- Informações Sensíveis nos Tokens: Tanto os tokens JWT quanto os tokens OAuth 2.0 contêm informações sobre o usuário, como ID, nome de usuário e quaisquer outros dados relevantes. Essas informações são criptografadas e podem ser transmitidas de forma segura entre o cliente e o servidor.

- Armazenamento no Cliente: Os tokens JWT e OAuth 2.0 são armazenados no cliente (geralmente em localStorage ou cookies) e enviados junto com cada solicitação subsequente ao servidor. Como os tokens são criptografados e assinados, eles não podem ser facilmente adulterados ou acessados por terceiros.

### Melhoria de Segurança

Além de garantir a confidencialidade das informações, a implementação de autenticação com JWT e OAuth 2.0 em uma aplicação em Go traz outras melhorias de segurança:

- Prevenção de Acesso Não Autorizado: Tanto JWT quanto OAuth 2.0 impedem o acesso não autorizado aos recursos protegidos da aplicação. Apenas usuários autenticados com tokens válidos têm permissão para acessar determinadas rotas e funcionalidades.

- Redução de Vulnerabilidades de Sessão: Ao usar tokens JWT e OAuth 2.0, a aplicação não precisa armazenar sessões no servidor, o que reduz a exposição a vulnerabilidades de gerenciamento de sessão.

- Maior Escalabilidade: A autenticação baseada em tokens JWT e OAuth 2.0 é escalável e independente de estado, permitindo que a aplicação atenda a um grande número de solicitações simultâneas de forma eficiente.

## JSON  Web Tokens (JWTs)

JSON Web Tokens (JWT) é um padrão aberto de mercado (RFC 7519) que define um formato compacto e autossuficiente para transmitir informações entre partes como um objeto JSON. Esse recurso é amplamente utilizado para autenticação e troca segura de informações entre sistemas.

### Métodos de Autenticação

Após um usuário realizar login com sucesso, o servidor gera um JWT que contém informações específicas sobre o usuário, como ID, nome de usuário, papel etc. Após sua geração, o token JWT é assinado digitalmente pelo servidor usando uma chave secreta (HS256, por exemplo), ou uma chave pública/privada (RS256), garantindo sua integridade. Além disso, O token JWT é enviado ao cliente e armazenado localmente (por exemplo, em localStorage ou cookies). Dessa forma, o cliente o envia junto com cada solicitação subsequente ao servidor.

### Complexidade de Implementação

A implementação básica de JWT é relativamente simples, exigindo apenas a geração e verificação de tokens. Muitas linguagens de programação, assim como 'Golang' utilizada no projeto, têm bibliotecas que facilitam esse processo. No entanto, o gerenciamento adequado de chaves para assinar e verificar tokens é crucial para garantir a segurança do sistema.

### Métodos e Funções associadas

Para o funcionamento da aplicação existem alguns métodos e funções associadas à sua implementação assim como a biblioteca "github.com/dgrijalva/jwt-go" em Go que é frequentemente usada para gerar, assinar e verificar tokens JWT. Dessa forma, pode-se implementar funções de "Verificação de Tokens", em que os servidores verificam a assinatura do token para garantir sua autenticidade e integridade antes de conceder acesso aos recursos protegidos e funções de "Decodificação de Tokens" em que os clientes podem decodificar o token JWT para acessar informações sobre o usuário sem a necessidade de consultar o servidor.

## OAuth 2.0

"OAuth 2.0" é um protocolo de autorização que permite que aplicativos obtenham acesso limitado a contas de usuário em um serviço HTTP. Ele é frequentemente utilizado para delegação de acesso, permitindo que um serviço acesse recursos protegidos em nome do usuário.

### Métodos de Autenticação

O OAuth 2.0 envolve vários fluxos de autorização, como o fluxo de autorização de código de autorização, o fluxo implícito e o fluxo de senha de proprietário do recurso. Além disso, com a implementação do fluxo de Consentimento do Usuário, o usuário deve conceder consentimento para que o cliente acesse seus recursos protegidos durante o processo de autorização.

### Complexidade de Implementação

A implementação de OAuth 2.0 pode ser mais simples em comparação com JWT devido à variedade de fluxos de autorização e à necessidade de interação com o servidor de autorização. É necessário configurar e gerenciar um servidor de autorização que lide com solicitações de autorização, emissão de tokens e consentimento do usuário.

### Métodos e Funções Associadas

Assim como o para o funcionamento da aplicação com a autenticação usando JWT, existem alguns métodos e funções associadas ao uso do OAuth 2.0 e várias linguagens de programação possuem bibliotécas para facilitar a implementação desta, como o "oauth2-proxy" em Go. 

Dessa maneira, para a implementação dessa tecnologia, os provedores de serviços OAuth exigem que os clientes sejam registrados por exemplo no Google, Facebook, GitHub e ou outras redes sociais e os fluxos de autorização devem ser implementados de acordo com as especificações desses provedores. Com essas autorizações, o cliente recebe um token de acesso, que pode ser usado para acessar os recursos protegidos em nome do usuário.

## Comparação JWT e OAuth 2.0

Para essa análise, pode-se analisar a implementação do ponto de vista da complexidade de implementação e escopo em que, enquanto o uso do JWT para autenticação é bem mais simples e direta entre dois sistemas confiáveis, o uso do OAuth exige um escopo mais amplo de autorização, especialmente em cenários envolvendo múltiplos serviços e interações com usuários finais, e portanto é de mais complexa implementação. Além disso, levando-se em consideração os pontos de segurança e complexidade, o JWT pode ser mais simples de implementar e gerenciar em comparação com OAuth 2.0, mas o OAuth 2.0 oferece uma estrutura mais robusta para autorização em sistemas complexos.

## Desenvolvimento inicial da aplicação com JWT

Para encotrar a implementação inicial da autenticação com  JWT, acesse: https://github.com/Inteli-College/2024-T0002-EC09-G04 e navegue até a pasta "backend/PD_auth". 

**Observação:** Nessa sprint, o desenvolvimento da autenticação ainda não foi integrada na aplicação principal do projeto. Esta, por sua vez, está sendo testada com dados fictícios a fim de compreender sua implementação e impacto no projeto.

### Funcionamento dos códigos:

- **`auth` Package:**

Este pacote contém uma função de middleware chamada TokenVerificationMiddleware que é responsável por verificar a presença e validade do token JWT nas requisições HTTP.

**Principais Funcionalidades:**

1. Verifica se o token JWT está presente no cabeçalho Authorization da requisição.
2. Valida o token JWT usando a chave secreta fornecida.
3. Caso o token seja inválido ou esteja ausente, retorna um status de não autorizado.

- **`handlers_auth` Package:**

Este pacote contém manipuladores de requisição para autenticação de usuários, incluindo registro, login e obtenção de lista de usuários.

**Principais Funcionalidades:**

1. `SignupHandler`: Manipulador para registrar novos usuários na aplicação.
2. `GetUsersHandler`: Manipulador para obter uma lista de todos os usuários registrados.
3. `LoginHandler`: Manipulador para autenticar usuários e gerar um token JWT válido para sessão.

### Principais Funções:

- **Middleware de Verificação de Token JWT**: A função TokenVerificationMiddleware é responsável por verificar a validade do token JWT nas requisições HTTP.
- **Geração de Token JWT**: A função generateToken é responsável por gerar um token JWT válido após a autenticação bem-sucedida de um usuário.
- **Rotas de Autenticação**: As rotas de autenticação fornecem endpoints para registro, login e obtenção de usuários. Elas são protegidas pelo middleware de verificação de token JWT, garantindo que apenas usuários autenticados possam acessá-las.

## Testes

**Teste `TestSignupHandler`:**

Este teste verifica o funcionamento do manipulador SignupHandler, que é responsável por registrar novos usuários na aplicação.

- **Funcionalidades Testadas:**
    - Cria um usuário de teste com nome de usuário e senha.
    - Converte o usuário em formato JSON para enviar na requisição HTTP.
    - Faz uma requisição POST ao endpoint `/signup` da API, passando o usuário como corpo da requisição.
    - Verifica se o código de status da resposta é `http.StatusOK`.
    - Verifica se o corpo da resposta contém o usuário registrado com sucesso.

**Teste `TestLoginHandler`:**

Este teste verifica o funcionamento do manipulador LoginHandler, responsável por autenticar os usuários e gerar tokens JWT válidos.

- **Funcionalidades Testadas:**
    - Cria um corpo de requisição com as credenciais de usuário (nome de usuário e senha) no formato JSON.
    - Faz uma requisição POST ao endpoint /login da API, passando as credenciais como corpo da requisição.
    - Verifica se o código de status da resposta é http.StatusOK.
    - Verifica se o tipo de conteúdo da resposta é application/json.
    - Verifica se há um token JWT válido na resposta.

## Conclusão

Em conclusão, a implementação de autenticação com JWT (JSON Web Tokens) e OAuth 2.0 em uma aplicação em Go oferece uma maneira segura e eficiente de garantir a confidencialidade das informações, além de melhorar a segurança geral do sistema.

Ao usar tokens JWT, podemos autenticar usuários de forma eficaz, garantindo que apenas usuários autorizados tenham acesso aos recursos protegidos da aplicação. Além disso, os tokens JWT permitem transmitir informações de forma segura entre o cliente e o servidor, contribuindo para a integridade e autenticidade dos dados.

Por outro lado, o OAuth 2.0 proporciona uma estrutura robusta para autorização, permitindo delegar acesso a recursos protegidos em nome dos usuários. Com o OAuth 2.0, podemos garantir que apenas aplicativos confiáveis tenham acesso aos dados do usuário, proporcionando uma camada adicional de segurança.





