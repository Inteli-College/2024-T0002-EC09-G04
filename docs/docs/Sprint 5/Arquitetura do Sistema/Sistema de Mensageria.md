---
title: Sistema de Mensageria 
sidebar_position: 8
slug: /mensageria5
---

# Conclusão sobre o Sistema de Mensageria com HiveMQ e Kafka

Neste documento, exploramos os conceitos e funcionalidades essenciais do sistema de mensageria utilizando HiveMQ em conjunto com Kafka por meio do Confluent Cloud, detalhando sua integração e funcionamento.

## Principais Pontos Abordados:

- **HiveMQ**: Uma plataforma de mensageria MQTT que fornece uma infraestrutura escalável e de baixa latência para troca de mensagens.
- **Kafka e Confluent Cloud**: Kafka é uma plataforma de streaming distribuída e o Confluent Cloud oferece uma solução gerenciada para Kafka, permitindo a criação, configuração e gerenciamento de clusters Kafka na nuvem.
- **Integração**: HiveMQ atua como um broker MQTT, enquanto o Kafka é usado para persistir e processar as mensagens, garantindo alta disponibilidade e confiabilidade.
- **Filas**: Utilizadas para armazenar mensagens de maneira temporária, operam sob o modelo FIFO, garantindo a ordem de entrega conforme o recebimento.
- **Publicação e Consumo de Mensagens**: Publicadores enviam mensagens que são roteadas para as filas apropriadas, enquanto os consumidores processam essas mensagens, podendo ajustar o consumo conforme a necessidade através do QoS.
- **Garantias de Entrega e Persistência**: Kafka assegura que as mensagens sejam entregues a pelo menos um consumidor e oferece a opção de persistência para proteção contra perdas em falhas.

## Funcionalidade e Gerenciamento:

A integração entre HiveMQ e Kafka proporciona uma solução robusta e escalável para sistemas de mensageria, oferecendo não apenas eficiência na troca de mensagens, mas também ferramentas poderosas para monitoramento e gerenciamento do ambiente de mensageria.

##  Rodando o Serviço por Meio de Interfaces 

Os serviços do Kafka, Confluent e HiveMQ são disponibilizados por meio de interfaces criadas com foco na facilidade de conexões e experiência do usuário. Entretanto a utilização desses serviços desse modo implica em restrições de utilização, como os _free tiers_ por exemplo. 

Pensando nisso, nosso desenvolvimento inicial foi aproveitando as limitações gratuitas de cada um deles, mas para o encerramento do projeto focamos em uma estrutura Cloud Agnostic Provider e verticalizamos a nossa infraestrutura e é possível rodar tudo o que estaria em serviços de nuvem localmente, o que posteriormente pode ser implantado em qualquer cloud provider.

## Rodando Localmente com Docker:

Para garantir uma opção de utilização mais agnóstica do projeto, é possível configurar e rodar a arquitetura de mensageria localmente utilizando containers Docker. Siga as instruções abaixo após entrar no diretório ```backend```:

```bash
make infra
```

## Conclusão:

A integração entre HiveMQ e Kafka por meio do Confluent Cloud oferece uma solução poderosa e flexível para sistemas de mensageria, permitindo alta disponibilidade, escalabilidade e eficiência na troca de mensagens. Além disso, a opção de rodar a arquitetura localmente proporciona uma alternativa versátil para desenvolvimento e testes, garantindo a portabilidade do projeto em diferentes ambientes.

Este documento visa fornecer uma compreensão clara e abrangente de como configurar, utilizar e testar a integração entre HiveMQ e Kafka, facilitando assim a implementação de sistemas de mensageria robustos e eficientes em diversos cenários de aplicativos.