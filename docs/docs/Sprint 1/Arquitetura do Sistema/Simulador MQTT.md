---
title: Simulador MQTT
sidebar_position: 4
slug: /simulator
---

O Simulador MQTT cumpre um papel fundamental no desenvolvimento da solução aqui apresentada, já que permite a simulação de uma grande massa de dados sendo transmitida através de tópicos MQTT, sem a necessidade de dispositivos físicos.

## Desafios

- Simular a variação de dados entre os sensores.
- Simular a disposição dos sensores e torno de uma aréa análoga a de uma cidade.
- Simular uma grande quantidade de sensores trasmitindo dados coletados simultaneamente.

## Proposta

Analisando, os desafios propostos, modelamos uma solução baseada em um novo objeto "estação" que é composto de dois sensores na versão atual (sensores de qualidade do ar e de radiação e lumunosidade). Essas estações são dispostas ao redor de uma área estabelecida no código a partir de coordenadas x e y que podem ser interpretadas como latitude e logitude. Os objetos sensores ao serem instaciados utilizam uma "seed" randômica para definir o set de dados, obviamente dentro do intervalo possível apontado pelo [TAPI](https://docs.google.com/document/d/15Z7xMHzdsVBHOvIR-BQjqVmZohKrPUJz/edit).

## Implementação

Para ter acesso a implementação da proposta anterior, acesse o [README.md](https://github.com/Inteli-College/2024-T0002-EC09-G04/tree/main/backend#backend) do projeto.

## Interagindo com a simulação

Todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L15C2-L15C7).

##### Comando:

```bash
make run
```

##### Output:

```shell
broker-1      | 1708301841: New client connected from 172.22.0.3:39954 as station-2805345190559206082 (p2, c1, k30).
broker-1      | 1708301841: New connection from 172.22.0.4:51230 on port 1891.
broker-1      | 1708301841: New client connected from 172.22.0.4:51230 as subscriber (p2, c1, k30).
api-test-1    | Received: {"location":"{\"latitude\":36.000000,\"longitude\":16.000000}","gas":"{\"CO2\":526.000000,\"CO\":7.000000,\"NO2\":343.000000,\"MP10\":404.000000,\"MP25\":179.000000}","rad_lum":"{\"ET\":267.000000,\"LI\":3.000000,\"SR\":241.000000}","timestamp":"2024-02-19 00:17:21.311444485 +0000 UTC m=+0.004634198"} from: /stations
```

:::tip NOTE
- Este comando está subindo todos os serviços presentes no arquivo compose.yml. São eles, o broker local, a simulação e a api-test que está sendo usada, por hora apenas para mostrar o log do que está sendo transmitido pela simulação.
:::