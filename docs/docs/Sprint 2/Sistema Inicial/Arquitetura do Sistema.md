---
title: Arquitetura do Sistema
sidebar_position: 1
slug: /arquitecture-system
---

# Arquitetura de Comunicação de Dados

## Componentes:

1. **Publisher MQTT:**
   - Responsável por enviar dados para o broker MQTT.
   - Configurado para se conectar ao broker HiveMQ.

2. **Broker MQTT (HiveMQ):**
   - Serviço na nuvem que gerencia a comunicação entre os publishers e os subscribers MQTT.
   - Recebe os dados dos publishers e os encaminha para os subscribers.

3. **Subscriber MQTT:**
   - Subscreve para receber dados do broker MQTT.
   - Configurado para receber dados do broker HiveMQ.
   - Capta os dados publicados pelos publishers.

4. **RabbitMQ:**
   - Sistema de mensageria que gerencia filas de mensagens.
   - Recebe dados do subscriber MQTT e os coloca em uma fila para processamento.

5. **Consumer RabbitMQ:**
   - Consome dados da fila gerenciada pelo RabbitMQ.
   - Processa os dados conforme necessário.
   - Envia os dados processados para o banco de dados RDS PostgreSQL.

6. **Banco de Dados RDS PostgreSQL:**
   - Banco de dados relacional hospedado na AWS RDS.
   - Recebe os dados processados do consumer RabbitMQ.
   - Armazena os dados de forma persistente para uso futuro.
   
![Arquitetura da Aplicação](../../../static/img/arq_sys.png)

## Fluxo de Dados:

1. Os publishers MQTT enviam dados para o broker HiveMQ.
2. O broker HiveMQ encaminha os dados para o subscriber MQTT.
3. O subscriber MQTT capta os dados e os passa para o RabbitMQ.
4. O RabbitMQ coloca os dados em uma fila para processamento.
5. O consumer RabbitMQ consome os dados da fila, os processa e os envia para o banco de dados RDS PostgreSQL para armazenamento.

## Configuração e Tecnologias Utilizadas:

- **MQTT (Message Queuing Telemetry Transport):** Protocolo de mensagens leve e eficiente para comunicação entre dispositivos.
- **HiveMQ:** Broker MQTT hospedado na nuvem, utilizado para roteamento e entrega de mensagens.
- **RabbitMQ:** Sistema de mensageria que suporta diferentes protocolos e padrões de comunicação, como AMQP (Advanced Message Queuing Protocol).
- **AWS RDS PostgreSQL:** Serviço de banco de dados relacional gerenciado na AWS (Amazon Web Services), que oferece escalabilidade, alta disponibilidade e segurança para aplicativos.
