---
title: Test-Driven Development
sidebar_position: 6
slug: /tdd-v2
---

A adoção do TDD (Desenvolvimento Orientado a Testes) é uma prática que oferece inúmeros benefícios no processo de desenvolvimento de software. Além de proporcionar feedback imediato sobre as novas funcionalidades, como mencionado anteriormente, o TDD promove a criação de um código mais limpo. Isso ocorre porque os desenvolvedores escrevem códigos simples, focados na passagem dos testes, resultando em uma base de código mais clara e fácil de manter. Essa abordagem não apenas acelera o ciclo de desenvolvimento, mas também contribui para a sustentabilidade e qualidade a longo prazo do software. Segue abaixo uma breve explicação de como esse projeto implementa isso e as instruções para comprovar a esteira de testes desenvolvida.

## Backend

### Rodar testes

Aqui, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L1).

#### Comando:

```shell
make tests
```

#### Output:

```shell
✔ Container simulation  Started                                                                                                      0.0s 
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/cmd/app    [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/cmd/simulation     [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/kafka       [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository  [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/web [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase   [no test files]
?       github.com/Inteli-College/2024-T0002-EC09-G04/backend/tools      [no test files]
=== RUN   TestNewAlert
--- PASS: TestNewAlert (0.00s)
=== RUN   TestNewLog
--- PASS: TestNewLog (0.00s)
=== RUN   TestEntropy
--- PASS: TestEntropy (0.00s)
=== RUN   TestNewSensor
--- PASS: TestNewSensor (0.00s)
=== RUN   TestNewSensorPayload
--- PASS: TestNewSensorPayload (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/Inteli-College/2024-T0002-EC09-G04/backend     0.005s  coverage: 100.0% of statements
=== RUN   TestMqttIntegration
2024/03/12 10:33:48 Selecting all sensors from the MongoDB collection sensors
--- PASS: TestMqttIntegration (152.76s)
PASS
coverage: [no statements]
ok      github.com/Inteli-College/2024-T0002-EC09-G04/backend/test       152.771s        coverage: [no statements]
```

:::note
- A esteira de testes definida aqui é responsável por testar as entitdade do ssistema e por realizar um teste de integração garantindo a entrega, e integridade e a frequência de envio dos dados, assim como o QoS da msg.
:::

### Rodar a visualização da cobertura de testes

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L31).

:::note
- Em uma cenário real, não precisaríamos de um arquivo como esse porque os sensores seriam criados por um frontend+backend que inseriria no banco de dados dos sensores.
:::

#### Comando:

```bash
make coverage 
```

#### Output:
![output_coverage](https://github.com/Inteli-College/2024-T0002-EC09-G04/assets/89201795/00c0ecc5-525f-4ccc-880f-272638852300)

:::note
- Este comando está criando, a partir do arquivo `coverage_sheet.md`, uma visualização da cobertura de testes nos principais arquivos Go.
:::

### Apêndice

Testes automatizados são essenciais no desenvolvimento de software, integrados com ferramentas de CI/CD para proporcionar uma abordagem consistente e eficiente. Essa prática permite identificar rapidamente problemas, reduzir erros e acelerar o ciclo de entrega, resultando em software de maior qualidade e lançamentos mais rápidos e confiáveis. Implementamos uma primeira versão dessa estrategia, para mais detalhes acesse o [workflow de testes](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/.github/workflows/tests.yml) do Github Actions.

## Frontend

Com o intuito de garantir a qualidade do produto final e a sua integridade, foram desenvolvidos testes para a interface do usuário, que permitem a validação da usabilidade do sistema da forma que ele foi projetado.

### Ferramentas Utilizadas

Para a realização dos testes, foi utilizado o Playwright, uma ferramenta de automação de testes de interface de usuário que permite a execução de testes em navegadores web. O Playwright é uma ferramenta de código aberto que oferece suporte a vários navegadores, como Google Chrome, Microsoft Edge e Mozilla Firefox. Dessa forma, é possível garantir que o sistema seja testado em diferentes ambientes, garantindo a sua compatibilidade e acessibilidade.

### Testes Realizados

Para essa sprint, foram desenvolvidos testes principalmente para a validação do sistema de login, signup e criação de um novo alerta. A seguir, são apresentados os testes realizados e os resultados obtidos.

![Testes Frontend](../../../static/img/frontend_test.png)

### Testes pendentes

Entretanto, ainda há testes que estão apresentando falhas, uma vez que o sistema ainda não apresenta as funcionalidades implementadas. Porém, de acordo com a metodologia TDD, os testes são desenvolvidos antes da implementação das funcionalidades, de forma que, após a implementação, os testes possam ser executados novamente para garantir que as funcionalidades foram implementadas corretamente.

### Conclusão

O frontend está apresentando o comportamento esperado até o presente momento, restando apenas a implementação das funcionalidades para que os testes possam ser executados novamente e validados. Com isso, é possível garantir que o sistema está funcionando corretamente e que as funcionalidades estão funcionando como planejadas.