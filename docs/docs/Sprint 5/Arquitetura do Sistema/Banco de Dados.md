---
title: Análise de Bancos de Dados
sidebar_position: 2
slug: /DB-conclusion
---

## Banco de Dados Relacional (PostgreSQL)

O sistema de monitoramento ambiental em São Paulo utiliza um [banco de dados relacional](https://inteli-college.github.io/2024-T0002-EC09-G04/database) PostgreSQL para armazenar informações sobre estações de monitoramento, dados de sensores, alertas da população e outras informações relevantes. A estrutura do banco de dados foi cuidadosamente projetada para atender às necessidades de armazenamento, recuperação e análise de dados ambientais, proporcionando uma base sólida para tomada de decisões e ações.

A integração com o sistema ocorre através do fluxo de dados dos sensores, que são transmitidos para um broker RabbitMQ e, em seguida, consumidos por um aplicativo que os armazena no banco de dados PostgreSQL. Isso garante a integridade, persistência e acessibilidade dos dados coletados pelos sensores e alertas da população, essenciais para a análise e tomada de decisões pelos usuários do sistema.

O modelo lógico do banco de dados oferece uma representação organizada e eficiente das informações coletadas, proporcionando uma estrutura clara para armazenamento e recuperação de dados. A relação do banco de dados com o restante do sistema é fundamental para garantir a integridade dos dados e a funcionalidade do sistema como um todo.

## Banco de Dados da Aplicação (Metabase)

O banco de dados da aplicação Metabase é utilizado para armazenar configurações, dados de usuários, logs, histórico de consultas e outras informações relacionadas ao funcionamento da própria aplicação. Ele desempenha um papel crucial na gestão e personalização da aplicação, fornecendo uma base sólida para operações como configuração de preferências do usuário, gestão de usuários e monitoramento de atividades.

Através do banco de dados da aplicação, o Metabase registra o histórico de consultas realizadas pelos usuários, armazena informações sobre os usuários registrados e gerencia configurações de conexão com fontes de dados. Isso permite uma interação eficiente e personalizada com a aplicação, contribuindo para uma experiência do usuário mais intuitiva e satisfatória.

## Banco de Dados NoSQL (MongoDB)

O [MongoDB](https://inteli-college.github.io/2024-T0002-EC09-G04/nosql-db) oferece uma solução flexível e escalável para projetos que demandam agilidade na manipulação de dados. Sua estrutura baseada em documentos JSON permite adaptar o esquema conforme necessário, enquanto sua capacidade de escala horizontal e desempenho rápido o tornam ideal para lidar com grandes volumes de dados e cargas de trabalho intensivas.

A implementação de repositórios para interagir com o banco de dados MongoDB permite criar, armazenar e recuperar informações sobre sensores, alertas e logs de forma eficiente. Isso simplifica a interação entre o programa e o MongoDB, fornecendo uma abstração para operações de criação, leitura, atualização e exclusão de dados.

Em conclusão, a análise dos bancos de dados PostgreSQL, Metabase e MongoDB demonstra a importância de escolher a tecnologia de armazenamento adequada para cada aplicação, levando em consideração requisitos de escalabilidade, desempenho, flexibilidade e facilidade de integração. Cada banco de dados desempenha um papel único no sistema de monitoramento ambiental em São Paulo, contribuindo para sua funcionalidade, eficiência e sucesso geral.