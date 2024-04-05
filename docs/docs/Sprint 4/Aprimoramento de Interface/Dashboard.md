---
title: Dashboard - Metabase
sidebar_position: 1
slug: /dash3
---

## Melhorias no Dashboard de Visualização de Dados

Este documento detalha a execução de testes de performance realizados durante a sprint 4, com ênfase em testes de carga para avaliar a capacidade da infraestrutura desenvolvida pelo grupo ao longo do módulo. O objetivo era verificar se a estrutura poderia suportar a demanda prevista de sensores. Inicialmente, estimamos a distribuição de sensores na área do município de São Paulo, dividindo-a em quilômetros quadrados e implantando quatro sensores por quarteirão, totalizando 6084 sensores. Esta quantia foi adotada como base para os testes subsequentes.

## Visualizações Aprimoradas

### Novas Visualizações de Gráficos

Pensando no aprimoramento da visualização da interface dos dados, foram criadas abas específicas para cada formato de visualização dos dados, com o objetivo de tornar a interpretação mais intuitiva para os usuários. Detalhes sobre cada aba:

1. **Região**:
   - Esta aba fornece uma visualização geográfica dos sensores, permitindo aos usuários visualizar a distribuição espacial dos mesmos.
   - Os usuários podem acompanhar a atividade dos sensores e alertas regionais inseridos pelos próprios usuários, o que facilita a compreensão dos padrões de ocorrência de eventos.
   - A interface permite a interação com o mapa, possibilitando a visualização de informações específicas de cada região e sensor.

2. **Gases**:
   - Nesta aba, os usuários têm acesso a visualizações detalhadas dos dados relacionados à medição de gases.
   - Gráficos e métricas são apresentados de forma clara, permitindo uma compreensão fácil dos níveis de poluentes atmosféricos e outras substâncias nocivas.
   - Os usuários podem comparar diferentes tipos de gases ao longo do tempo e identificar tendências e padrões relevantes.

3. **Luminosidade**:
   - Aqui, os usuários podem visualizar e acompanhar dados relacionados à medição de níveis de luminosidade em diferentes áreas.
   - Gráficos de tendências de luminosidade ao longo do dia ou ao longo do ano são disponibilizados, auxiliando na compreensão dos padrões de iluminação natural.
   - A interface oferece ferramentas de análise para identificar variações sazonais e ajustar a iluminação pública de acordo com as necessidades.

### Informações de Instrução e Orientação

Além da inclusão de novos gráficos, desenvolvemos um sistema de textos e orientações para fornecer informações adicionais aos usuários sobre a interpretação dos dados. Detalhes sobre esta funcionalidade:

- Cada tipo de visualização de dados é acompanhado de instruções claras sobre como interpretar os gráficos e métricas apresentadas.
- Explicações sobre o impacto dos dados no cotidiano dos cidadãos são fornecidas, ajudando os usuários a compreender a relevância das informações para suas vidas diárias.
- As orientações são apresentadas de forma intuitiva, integradas à interface para facilitar o acesso e compreensão por parte dos usuários.

### Demonstração 
<iframe width="560" height="315" src="https://www.youtube.com/embed/MALgdnRloEw?si=NasgkIKvrP7Cviy6" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>