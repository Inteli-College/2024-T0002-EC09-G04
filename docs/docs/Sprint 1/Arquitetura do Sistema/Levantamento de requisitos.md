---
title: Levantamento de Requisitos
sidebar_position: 2
slug: /requirements-gathering
---

Para garantir que o projeto atenda às necessidades do cliente, identificadas durante uma entrevista de levantamento de requisitos em sala de aula, os requisitos foram divididos em duas categorias principais: funcionais e não funcionais. Essa divisão foi motivada pela complexidade do projeto, que demanda tanto a implementação de funcionalidades específicas conforme os requisitos mínimos estabelecidos quanto o cumprimento de métricas de desempenho. Dentro dessas categorias, os requisitos funcionais descrevem o que o sistema deve realizar, enquanto os não funcionais estabelecem as métricas de desempenho a serem alcançadas pelos requisitos funcionais. Cada requisito foi ainda classificado como "obrigatório" ou "desejável", permitindo à equipe de desenvolvimento estabelecer uma ordem de prioridade para sua implementação.

## Requisitos funcionais


| ID - Título                     |  Descrição    | Categoria  |
| --------------------------------| ------------- | ---------  |
| RF1 - Coleta de Dados Ambientais| O sistema deve ser capaz de coletar dados de sensores. Incluindo, qualidade do ar, ruído urbano, radiação solar, luminosidade e umidade do solo.             |       Obrigatório     |
| RF2 - Dashboard de Visualização de Dados | Criação de dashboard intuitivo e acessível ao público e gestores públicos, exibindo dados e permitindo análises específicas por região.            |           Obrigatório    |            
| RF3 - Plataforma de Participação Cidadã | Criação de uma plataforma online que permite aos cidadãos acessar dados, relatar problemas ambientais e fornecer feedbacks. | Obrigatório |
| RF4 - Análise de Dados e Relatórios |  Criação de uma funcionalidade para análise de dados coletados, capaz de gerar relatórios sobre as condições ambientais e tendências ao longo do tempo. | Desejável |
| RF5 - Gestão de Usuários | Criação de cadastro na plataforma com diferentes níveis de acesso, servidor público, cidadão ou acadêmico. | Desejável |

## Requisitos Não Funcionais

| ID - Título                     |  Descrição    | Métrica    | Categoria  |
| --------------------------------| ------------- | ---------  | -----------|
| RNF1 - Coleta de Dados Ambientais | O sistema de coleta de dados deve garantir uma alta taxa de disponibilidade, recebendo ao menos 90% dos dados enviados,garantindo a precisão dos dados coletados. | Taxa de Transmissão de Pacotes (TTP) | Desejável |
| RNF2 - Dashboard de Visualização de Dados | O dashboard deve ser projetado com uma interface intuitiva, utilizando recursos que garantam acessibilidade e garantindo que usuários de todos os níveis técnicos possam entender facilmente as informações apresentadas. | Responsividade e testes com usuários | Obrigatório |
| RNF3 - Plataforma de Participação Cidadã | O sistema da plataforma de participação cidadã deve garantir uma disponibilidade contínua e confiabilidade elevada, com um tempo de atividade mínimo de 99,9% ao longo de cada mês civil. | Tempo de Indisponibilidade Programada e Não Programada | Obrigatório |
| RNF4 - Análise de Dados e Relatórios | As ferramentas de análise de dados devem processar grandes volumes de informações de forma eficiente, proporcionando análises e relatórios em tempo real sem atrasos significativos. | Tempo de processamento de dados | Obrigatório |
| RNF5 - Gestão de Usuários | O sistema deve garantir a segurança dos dados e o controle de acesso, permitindo diferentes níveis de acesso para servidores públicos, cidadãos e acadêmicos. | Taxa de Acesso não autorizado | Desejável |



