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

Para ter acesso a implementação da proposta anterior, acesse o [README.md](https://github.com/Inteli-College/2024-T0002-EC09-G04/tree/main/backend#backend) do projeto e siga as intruções para interagir com sistema.