---
title: Arquitetura da solução
sidebar_position: 1
slug: /solution-architecture
---
![Arquitetura - Versão Inicial](/img/arquitetura1.png)

## Camada de Segurança
Abriga todo o projeto, a proporção que se refere a todas as validações, redes, testes e outros métodos de segurança que serão implementados.

## Sistema de Geração de Informações
Responsável pela geração e processamento de dados sensoriais
- Dados Sensoriais
- Sensores de Qualidade do Ar (CO2, CO, NO2, MP10, MP2,5)
- Sensores de Ruído Urbano
- Sensores de Luminosidade
- Gerenciamento de Filas

## Sistema de Consumo dos Dados (MQTT)
Gerencia a troca de dados utilizando o protocolo MQTT
- Broker
- Publisher
- Subscriber
- Captação dos Dados (MQTT)

## Sistema de Serviço (Backend)
Responsável pelo servidor e pela lógica de negócios
- Servidor
- Armazenamento de Dados Sensoriais
- APIS
- Alimentação da Interface
- Conexão de Envio: Cloud
- Sistema de Testes Unitários

## Sistema de Armazenamento e Gestão de Dados em Cloud
Responsável pelo armazenamento e gerenciamento de dados na nuvem
- Banco de Dados não Relacional
- Serviço de Gerenciamento de Dados

## Interface de Interação do Usuário
Interface para interação do usuário com o sistema
- Sistema de Autenticação
- Níveis de Usuários
- Fórum de Contribuição da Comunidade
- Área de Informações e Orientações
