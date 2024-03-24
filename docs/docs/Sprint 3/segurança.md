---
title: Análise de Segurança
sidebar_position: 1
slug: /security
---

## Introdução

Com o intuito de nos fazer refletir sobre a segurança do projeto, foi proposta uma atividade em que deveríamos aprender mais sobre as técnicas de segurança e vulnerabilidades que podem ser encontradas em um sistema web. Para isso, nos foi apresentado uma atividade em que deveríamos explorar um sistema de login e identificar como podemos obter acesso a informações que não deveriam ser acessíveis a qualquer usuário.

## Ferramentas Utilizadas

Para a realização da atividade, foi explorado um sistema de login desenvolvido em PHP, que utiliza um banco de dados MySQL para armazenar as informações dos usuários, e para o deploy, foi utilizado o NGINX.

## Testes realizados

### Primeiro Teste

Para o primeiro teste, foi utilizado um método de injeção de SQL que realiza o dump de todos os usuários, apresentando a informação dos seus salários. Para isso, foi inputado o seguinte valor no campo de usuário:

```sql
' OR 1=1#
```

Os detalhes de funcionamento desse comando estão detalhados a seguir:

* ' OR: Aqui, a aspa simples ('), seguida por "OR", sugere uma inserção de condição lógica em uma consulta SQL. O "OR" é um operador lógico usado para combinar duas ou mais condições em uma expressão. Nesse contexto, sugere-se a adição de uma condição adicional à consulta SQL.

* 1=1: Esta é uma expressão lógica verdadeira em SQL. Quando 1 é igual a 1, a expressão é sempre verdadeira. Portanto, isso está sendo usado para garantir que a condição adicionada com o operador "OR" seja sempre verdadeira.

* #: O símbolo # é frequentemente usado para iniciar um comentário em SQL. Qualquer coisa depois disso será ignorada pelo interpretador SQL. Isso é útil para encerrar qualquer consulta SQL que esteja sendo injetada no sistema e evitar erros de sintaxe.

O resultado obtido foi a listagem de todos os usuários e seus salários, como pode ser visto na imagem a seguir:

![security1](../../static/img/security1.png)
![security2](../../static/img/security2.png)

### Segundo Teste

Para o segundo teste, foi utilizado um método de injeção de SQL que obtém a senha de todos os usuários cadastrados no sistema. Para isso, foi inputado o seguinte valor no campo de usuário:

```sql
' UNION SELECT username,password FROM users#
```

Os detalhes de funcionamento desse comando estão detalhados a seguir:

* UNION SELECT: A cláusula UNION em SQL é usada para combinar os resultados de duas ou mais consultas SELECT em uma única tabela de resultados. Esta parte do comando sugere uma união do resultado de uma consulta legítima com outro conjunto de dados que eles estão tentando extrair. O SELECT indica que eles querem selecionar certas colunas de dados.

* username,password: Aqui, estão sendo especificadas as colunas que desejam selecionar, visando obter os nomes de usuário e senhas armazenadas no banco de dados.

* FROM users: Esta parte indica a seleção dos dados da tabela chamada "users", onde estão armazenadas as informações de usuário e senha.

* #: Assim como no comando anterior, o símbolo # é usado para iniciar um comentário em SQL, o que significa que tudo após ele será ignorado pelo interpretador SQL. Isso é feito para encerrar a consulta SQL e evitar erros de sintaxe.

![security3](../../static/img/security3.png)
![security4](../../static/img/security4.png)

### Utilização da ferramenta SQLMap

Além dos testes realizados de forma manual, também foi utilizado a ferramenta SQLMap para realizar testes de injeção de SQL. O SQLMap é uma ferramenta de teste de penetração que automatiza o processo de detecção e exploração de falhas de injeção de SQL. Ele é capaz de detectar e explorar várias falhas de injeção de SQL, como injeção de SQL cega, injeção de SQL baseada em erro, injeção de SQL baseada em tempo, injeção de SQL booleana, entre outras.

Para realizar o teste, foi utilizado o seguinte comando:

```bash
python sqlmap.py -u http://localhost:8080/ --batch --banner --method POST --forms --dump
```

O resultado obtido foi a listagem de todos os usuários e suas senhas, como pode ser visto na imagem a seguir:

![security5](../../static/img/security5.png)

### Conclusão

A atividade nos permitiu entender a importância de se preocupar com a segurança de um sistema, e como é importante realizar testes de segurança para garantir que o sistema não esteja vulnerável a ataques. Através dos testes realizados, foi possível observar o quão fácil é explorar um sistema desprotegido, e como devemos focar nossos esforços em garantir que o sistema esteja seguro.

## Possíveis vulnerabilidades e ataques ao sistema

No caso de uma autenticação com senhas simples ou com padrões facilmente decifráveis, o sistema fica exposto à vulnerabilidades como ataques que ferem a confidencialidade, disponbilidade ou integridade do sistema. A fim de mitigar possíveis falhas ou situações que comprometem a aplicação, métodos como
