---
title: Analise de Acessibilidade - Guia WCAG
sidebar_position: 2
slug: /acessibility
---

## Aplicação do Guia WCAG

As WCAG, ou Web Content Accessibility Guidelines (Diretrizes de Acessibilidade para Conteúdo da Web), constituem um conjunto de diretrizes e recomendações estabelecidas e mantidas pelo W3C (World Wide Web Consortium), visando orientar a criação de conteúdos digitais que sejam acessíveis a todas as pessoas, independentemente de suas habilidades ou necessidades.

Tendo isso em mente, ao desenvolvermos um projeto voltado para a criação de uma interface destinada ao acompanhamento e visualização de dados em uma cidade inteligente, é crucial considerarmos a acessibilidade como um aspecto primordial. Nesse contexto, nosso objetivo é conduzir uma análise abrangente visando aprimorar os aspectos relacionados à acessibilidade, reconhecendo que as cidades inteligentes não se limitam apenas ao desenvolvimento tecnológico, mas também incorporam uma abordagem inteligente e inovadora em relação à inclusão social.

## Objetivo

Nosso objetivo consiste na seleção de um card em cada uma das categorias: Operável, Perceptível e Compreensível, do <a href="https://guia-wcag.com/" target="_blank">Guia WCAG</a>. Com base nisso, realizamos um levantamento sobre os cards selecionados e desenvolvemos uma explicação  das funcionalidades de cada um em relação às interfaces existentes no projeto para a interação do usuário, conforme o projeto que está sendo desenvolvido.

## Interfaces do Projeto

1. Dashboard
Um dashboard permite que usuários explorem, visualizem e compartilhem insights a partir de conjuntos de dados. No caso específico do projeto, com a utilização do Metabase não há a necessidade de conhecimento avançado em SQL ou programação. O Metabase facilita a criação de painéis interativos e consultas personalizadas, tornando a análise de dados mais acessível para uma variedade de usuários. Nesse projeto, ele é utilizado para a visualização das entidades presentes no banco de dados, criação de gráficos, mapas e outros insights. 

2. Entrada de Dados do Usuário: 
Para além do dashboard de visualização dos dados, foi desenvolvida uma interface de front-end que consiste na área que permite a inclusão de novos dados de alerta e sensores. Nesse sistema temos algumas funções principais, um sistema de login que permite que os usuários autentiquem-se para acessar as funcionalidades do sistema, enquanto os formulários de envio de dados permitem que os usuários enviem alertas e a criação de novos sensores.

O frontend oferece duas principais funcionalidades:

[x] - Criação de Estações de Sensores: Os gestores podem criar estações com os sensores necessários para medir determinadas variáveis em uma região específica. Os dados são enviados em um formato específico, incluindo nome, latitude, longitude e parâmetros de cada sensor.

[x] -Alertas de Incidentes: Os usuários comuns podem enviar alertas relatando incidentes encontrados durante suas atividades. Os dados são enviados em um formato simples, incluindo latitude, longitude e descrição do incidente.


## Classificações WCAG 
### Níveis de Conformidade

#### A
A geralmente significa que o conteúdo da web será acessível para algumas pessoas com necessidades, mas não necessariamente para todas.

#### AA
Significa que o conteúdo da web será acessível para a maioria das pessoas com necessidades.

#### AAA
Significa que o conteúdo da web será acessível para quase todas as pessoas com necessidades.

## Seleção de Cards para Aplicação ao Projeto

### 1.4.6 - Contraste (Melhorado) [AAA]
### _Perceptível_

1. Exceções e flexibilidade: 
Elementos decorativos são excluídos das diretrizes; textos maiores têm requisitos de contraste mais flexíveis para permitir liberdade criativa aos designers.

2. Aplicação prática: 
Dashboard:  
Teste de Atualização de Cores do Dash para relações de cores com maior contraste. 
Aplicação Web: 
Focar em necessidades de ação e interação do usuário com cores de maior contraste para que o usuário possa utilizar o contraste como fator de localização e guia do fluxo das ações.

### 3.2.3 - Navegação Consistente [AA]

### _Compreensível_

1. Importância da Navegação Consistente:
Essencial para facilitar a navegação eficiente, especialmente para usuários com baixa visão ou que usam ampliação de tela.

2. Aplicação na Interface:
No dashboard e na entrada de dados do usuário, a consistência na apresentação de informações e funcionalidades é crucial.
Garantir que a informação e os layouts sigam um fluxo mais constante. Ou seja, o usuário (qualquer usuário) possa “reconhecer” sempre qual é o próximo passo que ele deve dar.

### 2.2.5 - Nova Autenticação [AAA]

### _Operável_

Garante que os usuários possam retomar suas tarefas sem perder progresso ou dados ao fazer login novamente.

2. Impactos Positivos no Dashboard:
Oferece uma experiência contínua aos usuários, especialmente em ambientes de monitoramento em tempo real, evitando perda de dados e decisões inadequadas.

Beneficia usuários com deficiência visual, eliminando a necessidade de inserir novamente dados ou reiniciar tarefas após login expirado.

3. Medidas de Melhoria:
Ajustar configurações de tempo de expiração da sessão para atender às necessidades dos usuários e contexto de uso.