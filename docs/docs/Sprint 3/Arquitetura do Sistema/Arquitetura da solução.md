---
title: Arquitetura da solução
sidebar_position: 1
slug: /solution-architecture-v3
---

![Arquitetura - V3](/img/arquitetura2.png)

## Camada de Segurança
Abriga todo o projeto, a proporção que se refere a todas as validações, redes, testes e outros métodos de segurança que serão implementados.

## Simulação
Gerencia a troca de dados utilizando o protocolo MQTT
- Sensores de Qualidade do Ar (CO2, CO, NO2, MP10, MP2,5);
- Sensores de Luminosidade e radiação;
- Broker ( HiveMQ Cloud );

## APP
Responsável pelo servidor e pela lógica de negócios
- Mensageria Kafka ( Confluent Cloud );
- APIS para criação de alertas e sensores
- Consumer Kafka
- Armazenamento de Dados Sensoriais

### Domain Driven Development & Arquitetura Hexagonal
Domain-Driven Development (DDD) combinado com a arquitetura hexagonal é uma abordagem de desenvolvimento de software que enfatiza a modelagem do domínio em um contexto de negócios, juntamente com uma arquitetura modular e desacoplada. No DDD, especialistas do domínio colaboram com desenvolvedores para criar um modelo de domínio claro e preciso, enquanto na arquitetura hexagonal, o sistema é dividido em camadas independentes que se comunicam por meio de interfaces bem definidas. Essa abordagem facilita a separação de preocupações e a manutenção do código, permitindo que a lógica de negócios seja isolada do código de infraestrutura. Ao combinar DDD e arquitetura hexagonal, o software resultante é altamente alinhado com os requisitos do negócio, fácil de entender e modificar, e permite uma evolução sustentável ao longo do tempo, adaptando-se facilmente a mudanças nos requisitos do domínio. Neste recorte do projeto, foi implentada essa estratégia pensando na manutenção do sistema, gerenciamento do sistema, escalabilidade e desacoplamento das regras de negócio e os seus respectivos adaptadores.

## Dashboard e Visualização de dados
Responsável pela visualização de gráfica dos dados coletados pela simulação.
- Banco de dados relacional ( AWS RDS );
- Metabase;

## Interface de Interação do Usuário
Interface para interação do usuário com o sistema
- Sistema de Autenticação
- Níveis de Usuários
- Fórum de Contribuição da Comunidade
- Área de Informações e Orientações

## Diagrama UML

![UML Diagram (1)](https://github.com/Inteli-College/2024-T0002-EC09-G04/assets/99187952/6337aac8-bb48-408b-92ca-4c19cfcaa4f2)

## Fluxo de Dados:

1. O Sistema de Simulação de Sensores gera dados dos sensores.
2. Os dados são consumidos pelo Sistema de Consumo MQTT, que os encaminha para o Sistema de Mensageria.
3. O Sistema de Mensageria gerencia as filas de dados utilizando Kafka e replica os dados no banco de dados.
4. O Sistema de Serviço Backend consome os dados do banco, os processa e os armazena no MongoDB.
5. As interfaces de visualização e interação acessam os dados armazenados para apresentação ao usuário.
