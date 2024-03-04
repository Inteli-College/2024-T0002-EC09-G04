---
title: Mensageria
sidebar_position: 3
slug: /menssager
---

Um dos componentes do back-end do sistema é caracterizado por agir como um consumidor dos dados produzidos. Uma vez que os dados recebidos pelo broker são enfileirados pelo RabbitMQ, o consumer estabelece conexão com o RabbitMQ, consome os dados da fila, desenfileirando-os e, por fim, popula o banco de dados PostgreSQL instanciado no RDS.

### Arquitetura e Fluxo de Dados

1. **Recebimento de Mensagens MQTT**: Inicializa uma conexão MQTT e se inscreve em um tópico específico para receber mensagens de dados de estações.
2. **Envio para RabbitMQ**: As mensagens MQTT recebidas são enviadas para uma fila no RabbitMQ.
3. **Consumo e Processamento de Mensagens**: O aplicativo consome mensagens da fila RabbitMQ, processa os dados JSON e insere os registros no banco de dados PostgreSQL.


### Funções

- `messagePubHandler`: Função para lidar com mensagens MQTT recebidas.
- `failOnError`: Auxiliar para tratamento de erros.
- `insertGasData`: Insere dados no banco de dados PostgreSQL.
- `sendToRabbitMQ`: Envia mensagens recebidas para uma fila no RabbitMQ.
- `main`: Inicializa a conexão MQTT, configura o tratamento de mensagens e se inscreve em um tópico.