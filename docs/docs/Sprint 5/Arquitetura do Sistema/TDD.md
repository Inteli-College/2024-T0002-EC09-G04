---
title: Test Driven Development
sidebar_position: 9
slug: /tdd5
---

## Testes de Carga

Este documento apresenta as conclusões dos [testes de desempenho](https://inteli-college.github.io/2024-T0002-EC09-G04/carga) realizados durante as sprints 4 e 5, com foco em testes de carga para avaliar a capacidade da infraestrutura desenvolvida pelo grupo ao longo do módulo. O principal objetivo era verificar se a estrutura seria capaz de suportar a demanda prevista de sensores. No início da sprint 4, o grupo fez uma estimativa inicial para compreender a magnitude do desafio. Optamos, posteriormente, por dividir a área do município de São Paulo em quilômetros quadrados e distribuir os sensores pelos quarteirões, o que resultou na decisão de implantar quatro sensores por quarteirão, totalizando 6084 sensores. Esta quantidade foi adotada como base para os testes subsequentes.

A escolha de uma arquitetura totalmente baseada em nuvem, abrangendo desde o broker até o sistema de mensageria, viabilizou a possibilidade de escalar o projeto horizontalmente e alavancar a distribuição do sistema. No entanto, essa decisão ocasionou algumas limitações, especialmente para a realização de testes, devido ao uso de planos gratuitos nos serviços de nuvem, que restringem o número de requisições. O entrave em questão nos levou a um ponto de estrangulamento, em que ultrapassar um determinado limite de requisições resultava na inutilização da conta para fins de teste.

Em resposta a essa limitação, consideramos a possibilidade de criar um cluster totalmente dockerizado da arquitetura. No entanto, a falta de familiaridade com a tecnologia de testes e o tempo restrito impediram a implementação dessa solução, levando a equipe a concentrar os testes dentro dos limites permitidos pela versão gratuita do HiveMQ, que suporta até 100 requisições simultâneas.

### Utilização do K6 para Testes de Carga

Diante do desafio de realizar nosso primeiro teste de carga, a equipe optou pelo K6, um framework baseado em Go que executa scripts de teste escritos em Javascript e suporta integração com o Grafana para monitoramento. O processo de teste envolve a definição de variáveis pré-estabelecidas e de ambiente, permitindo testes tanto em ambientes locais quanto em nuvem. Para nossos testes em nuvem, configuramos variáveis como host, usuário e senha. Ao executar o script, o K6 gera várias instâncias do script para simular o número desejado de clientes, publishers ou subscribers, alcançando o limite de 100 conexões imposto pelo Free Tier do HiveMQ. Ao concluir a execução, o script fornece um relatório no console com os resultados do teste, incluindo o tempo de execução e o número de clientes simulados.

Durante a quinta sprint, em uma tentativa de utilizar o Grafana para integração com o K6 e visualização de testes, o grupo descobriu duas possibilidades de contorno para as limitações da arquitetura: executar a arquitetura do sistema em ambiente local, com o HiveMQ e Kafka como contêineres Docker, e exibir os resultados dos testes no console, sem demonstração de gráficos. Por outro lado, se o grupo quisesse exibir esses resultados com a integração do Grafana, a forma de fazer isso poderia ser subir uma imagem do Grafana em um EC2 AWS, o que contornaria a limitação de 100 clientes imposta pelo Grafana.

Em suma, a implementação desse teste de carga durante as duas últimas sprints teve foco na descoberta das melhores práticas de teste e das ferramentas utilizadas para realizá-lo. Isso se mostrou promissor para sprints que envolvem serviços de servidores para os próximos módulos e refinamento de arquiteturas que podem estar um pouco lentas ou sobrecarregadas.

## Testes de Autenticação

Com o intuito de apresentar uma solução segura e eficiente, o grupo realizou testes de autenticação para garantir a integridade do sistema. A autenticação é um processo fundamental para a segurança de sistemas, pois permite a verificação da identidade de usuários e dispositivos, garantindo que apenas usuários autorizados tenham acesso ao sistema. Para isso, o grupo implementou um sistema de autenticação baseado em JWT (JSON Web Token), que permite a autenticação de usuários por meio de tokens criptografados.

O teste da autenticação foi realizado por meio de um script de teste automatizado, que simula o processo de autenticação de usuários no sistema. O script envia uma solicitação de autenticação para o servidor, que verifica as credenciais do usuário e retorna um token JWT válido. O token é então utilizado para autenticar o usuário em solicitações subsequentes, garantindo que apenas usuários autenticados tenham acesso aos recursos protegidos do sistema.

Para isso, foi utilizado o framework de testes automatizados Playwright, que permite a automação de testes de interface de usuário em navegadores web. O script de teste foi desenvolvido para simular o processo de autenticação de usuários no sistema, verificando se o token JWT gerado pelo servidor é válido e se o usuário é autenticado corretamente. O script foi executado em um ambiente de teste isolado, garantindo a integridade dos resultados.

Os testes estão disponíveis no repositório do grupo, no diretório `tests`, presentes na pasta `frontend`, podendo ser acessados pelo seguinte [link](https://github.com/Inteli-College/2024-T0002-EC09-G04/tree/main/frontend/tests)
