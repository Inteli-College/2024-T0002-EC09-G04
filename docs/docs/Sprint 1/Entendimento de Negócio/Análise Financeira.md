---
title: Análise Financeira
sidebar_position: 2
slug: /financial-analysis
---

Neste relatório, apresentaremos uma análise financeira detalhada do projeto. O sistema será analisado em diferentes fases, desde o desenvolvimento até a análise de receita futura, incluindo os custos associados a cada etapa. Com isso, objetivamos fornecer um panorama completo do investimento necessário para a implementação da solução proposta.

## Fases analisadas

- **Desenvolvimento:** Custo de pesquisa, design e programação.
- **Produção:** Custo de produtos, materiais, fabricação e montagem.
- **Operação:** Custo de manutenção, energia e possíveis atualizações.
- **Receita:** Montante total de dinheiro gerado pela venda da solução.

#### Referenciais:

  - A área do Estado de São Paulo corresponde a 248,219 km².
  - A média salarial de um Engenheiro Elétrico é R$ 8.253,10.
  - A média salarial de um Engenheiro de Computação é R$ 8.479,39.
  - A média salarial de um Engenheiro de Software é R$ 7.000,00.
  - A média salarial de um técnico em TI é R$ 1.729,00.
  - A Bolsa-auxílio média para estágio em Engenharia de Computação é R$ 1.157,00.
  - A Bolsa-auxílio média para estágio em Engenharia Elétrica é R$ 1.278,00.
  - Edital para licitações do Finep/MCTI datado de 2022.
  - Contratações realizadas pelo IPT - Instituto de Pesquisas Tecnológicas.
  - A fabricação dos componentes definidos na etapa de pesquisas será nacionalizada.
  - O tempo para a execução do projeto deverá ser de 12 meses.

## Desenvolvimento

Analisando o projeto como uma contratação terceirizada para desenvolvimento de software e implementação da infraestrutura física, podemos inferir o custo associado à mão de obra, definindo que a cada 200 km² teremos 2 engenheiros chefes, 4 estagiários e 10 técnicos. Adotando uma distribuição/demanda de mão de obra que segue graficamente uma distribuição normal padrão e adicionando no valor final custos adjuntos à operação, como o deslocamento das equipes, chegamos a um valor desde o início, passando pelo planejamento da execução do projeto, operações em campo para a instalação da infraestrutura, até a finalização do projeto de:

- **Mão de obra:** R$ 1.792.664,34;
- **Custos adjacentes:** R$ 450.000,00.

**Custo total estimado: R$ 2.242.664,34.**

## Produção

Levando em consideração as etapas estabelecidas pela grande maioria dos frameworks de execução de projetos modernos, colocaremos a etapa de P&D (Pesquisa e Desenvolvimento) como essencial na análise dos custos de produção, seguida pelo custo dos componentes físicos e do software requerido.

Para a fabricação do software, vamos estimar um time de 25 Engenheiros de Software durante 8 meses, também seguindo uma distribuição normal padrão.

- **P&D:** R$ 1.800.000,00;
- **Fabricação dos Componentes:** R$ 2.600.000,00;
- **Fabricação do Software:** R$ 700.000,00.

:::note
A mão de obra dessa fase foi ocultada porque o proposto nessa análise é a contratação de um instituto de pesquisas e, portanto, a delegação para esse da contratação dos profissionais necessários.
:::

**Custo total estimado: R$ 5.100.000,00.**

## Operação

Tendo em vista a manutenção de tudo o que será implementado, é certo que os custos não cessarão na entrega do projeto. Nesse sentido, podemos estimar duas principais verticais (os custos abaixo retratados são de caráter mensal):

- **Infraestrutura Cloud para o sistema:**
  - Serviço Cloud: R$ 15.000,00.
  - Mão de obra: R$ 42.000,00.
- **Manutenção da infraestrutura física:**
  - Peças sobressalentes: R$ 26.000,00 (4% do total de peças necessitará de manutenção a cada 2 meses).
  - Mão de obra: R$ 43.079,39.
  - Custos adjacentes: R$ 8.645,00.

#### Detalhamento:

- Sobre a infraestrutura cloud, temos que levar em consideração que o sistema é altamente distribuído e que ele terá que lidar com um grande tráfego de dados. O valor apontado se baseia nos preços da maior infraestrutura cloud de SP na AWS ([acesse aqui](https://aws.amazon.com/pt/pricing/?aws-products-pricing.sort-by=item.additionalFields.productNameLowercase&aws-products-pricing.sort-order=asc&awsf.Free%20Tier%20Type=*all&awsf.tech-category=*all)), o que garante uma latência menor do sistema. Para gerenciar e realizar a manutenção desse sistema, também será necessário pelo menos 6 Engenheiros de Software por mês enquanto o sistema estiver ativo.

- Em relação à manutenção da infraestrutura física, o principal motivo será danificações decorridas de intempéries naturais e troca de baterias. Para realizar tais manutenções, será necessário pelo menos 15 técnicos e 1 Engenheiro Chefe por mês realizando operações em campo enquanto o sistema estiver ativo.

**Custo total estimado: R$ 134.724,39.**

## Receita
É importante destacar que o modelo de negócios aqui proposto, por ter ligação com entes públicos, não tem como objetivo uma receita direta advinda do projeto aqui desenvolvido. Contudo, isso não descarta a possibilidade desse projeto gerar uma maior eficiência no controle de gastos públicos e até mesmo na diminuição deles, à medida que fornece dados que podem ser utilizados para prever obras públicas ou até no gerenciamento de políticas públicas, como aquelas ligadas à saúde. Nesse sentido, o resultado final dessa balança, entre custo e resultado financeiro, pode gerar um resultado positivo, comparado a uma realidade em que, na ausência de uma solução como a proposta aqui, é necessário muito mais investimento por parte dos cofres públicos.

:::note Resultado
***Custo total de implementação: R$ 7.342.664,34***

*Diluído: R$ 611.888,70/mês durante 12 meses, que é o prazo estabelecido como referencial para a entrega final do projeto.*

***Custo total em três anos de operação: R$ 7.612.113,12***

*Diluído: R$ 293.779,16/mês.*
:::