---
title: Arquitetura do Sistema
sidebar_position: 1
slug: /arquitetura5
---

## Conclusão do Projeto

Este projeto abrange uma série de componentes interconectados que juntos formam uma solução robusta para gestão e análise de dados ambientais e de interação do usuário. A seguir, detalhamos as conclusões de cada componente fundamental do sistema.

### Camada de Segurança
Essencial para a proteção integral do projeto, essa camada compreende todas as validações necessárias, gestão de redes e testes, assim como outros métodos de segurança que são fundamentais para a integridade e confiabilidade do sistema.

### Simulação
Utiliza o protocolo MQTT para gerenciar a troca de dados entre diferentes sensores, incluindo sensores de qualidade do ar e luminosidade. O uso do broker HiveMQ Cloud assegura que o tráfego de dados seja eficiente e seguro.

### Aplicativo (APP)
Serve como o núcleo para a lógica de negócios e a operação do servidor, empregando Kafka para mensageria através da Confluent Cloud e oferecendo APIs para criação de alertas e gestão de sensores. Além disso, é responsável pelo armazenamento de dados sensoriais, consolidando a base para análise e decisão.

### Desenvolvimento Orientado ao Domínio & Arquitetura Hexagonal
A adoção do Domain-Driven Development (DDD) em conjunto com a arquitetura hexagonal promove uma estruturação clara e modular do software, facilitando a manutenção, escalabilidade e a separação entre as regras de negócio e a infraestrutura.

### Dashboard e Visualização de Dados
Responsável pela apresentação gráfica dos dados coletados, utilizando o AWS RDS como banco de dados relacional e o Metabase para a visualização e análise de dados, garantindo que as informações sejam facilmente acessíveis e compreensíveis.

### Interface de Interação do Usuário
Proporciona uma interação fluida e segura para os usuários finais do sistema, incorporando funcionalidades como sistema de autenticação, definição de níveis de usuário, além de fóruns e áreas de informações que enriquecem a experiência do usuário.

### Fluxo de Dados
O sistema inicia com a simulação de sensores que geram dados ambientais. Esses dados são então capturados pelo sistema MQTT, enviados para o sistema de mensageria Kafka e posteriormente armazenados em bancos de dados. O backend processa e disponibiliza esses dados para as interfaces de usuário, onde são visualizados e gerenciados.

Em resumo, o projeto integra tecnologias avançadas e metodologias de desenvolvimento modernas para criar uma solução eficaz que não apenas atende às necessidades de monitoramento ambiental, mas também proporciona uma interação rica e segura para os usuários, demonstrando a flexibilidade e a capacidade do sistema de adaptar-se a variadas demandas e escalas de uso.