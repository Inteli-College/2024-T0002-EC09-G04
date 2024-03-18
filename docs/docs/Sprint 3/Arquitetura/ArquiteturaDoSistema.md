---
title: Arquitetura do Sistema
sidebar_position: 1
slug: /arquitecture-system
---

# Arquitetura de Comunicação de Dados

## Componentes:

1. **Sistema de Simulação de Sensores:**
   - Responsável por simular dados dos sensores de qualidade do ar, luminosidade e radiação solar.
   - Utiliza a tecnologia Docker para encapsulamento e distribuição de aplicativos.

2. **Sistema de Consumo MQTT:**
   - Broker MQTT responsável por receber e encaminhar dados.
   - Subscriber MQTT para captação de dados.
   - Integração entre broker e fila para gerenciamento de mensagens.
  

3. **Sistema de Mensageria:**
   - Utiliza Kafka para gerenciar filas e replicação de banco de dados.
   - Integração entre broker e fila para comunicação eficiente.
   

4. **Sistema de Serviço Backend:**
   - Servidor para processamento de dados.
   - Consumer Kafka para consumo de dados da fila.
   - APIs para manipulação e processamento dos dados.
   - Conexão com MongoDB Atlas para armazenamento.
  

5. **Sistema de Armazenamento de Dados do Sistema:**
   - Banco de dados não relacional para armazenamento de logs da simulação, alertas e dados dos sensores.
   

6. **Sistema de Armazenamento de Dados do Metabase:**
   - Banco de dados relacional para armazenamento de dados da aplicação.
   - Serviço para persistência de dados.
   

7. **Interface de Visualização de Dados:**
   - Interface de visualização e business intelligence.
   

8. **Interface de Interação do Usuário:**
   - Dashboard interativo para níveis de usuários.
   - Fórum de contribuição da comunidade.
   - Área de informações e orientações.
   
![Frame 4](https://github.com/Inteli-College/2024-T0002-EC09-G04/assets/99187952/35e7cff1-4291-4d2c-bf76-d681f803a041)


## Fluxo de Dados:

1. O Sistema de Simulação de Sensores gera dados dos sensores.
2. Os dados são consumidos pelo Sistema de Consumo MQTT, que os encaminha para o Sistema de Mensageria.
3. O Sistema de Mensageria gerencia as filas de dados utilizando Kafka e replica os dados no banco de dados.
4. O Sistema de Serviço Backend consome os dados do banco, os processa e os armazena no MongoDB.
5. As interfaces de visualização e interação acessam os dados armazenados para apresentação ao usuário.

## Configuração e Tecnologias Utilizadas:

- **Docker:** Tecnologia de contêineres para distribuição de aplicativos.
- **MQTT (Message Queuing Telemetry Transport):** Protocolo de mensagens para comunicação entre dispositivos.
- **HiveMQ Cloud:** Broker MQTT na nuvem para roteamento de mensagens.
- **Kafka:** Sistema de mensageria para gerenciamento de filas e replicação de dados.
- **MongoDB:** Banco de dados não relacional para armazenamento de dados.
- **MongoDB Atlas:** Serviço de banco de dados MongoDB hospedado na nuvem.
- **AWS:** Amazon Web Services para hospedagem de serviços e armazenamento de dados.
- **Go:** Linguagem de programação utilizada no servidor backend.
- **Metabase:** Ferramenta de business intelligence para análise e visualização de dados.
