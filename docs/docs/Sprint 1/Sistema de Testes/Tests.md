---
title: Planejamento de Testes Unitários
sidebar_position: 1
slug: /unitary_tests
---

# Sistemas de Testes do Projeto

Neste documento, vamos detalhar o planejamento de testes unitários para o projeto EcoVigilância, focando no sistema de broker MQTT, no servidor e na segurança das aplicações. Os testes serão desenvolvidos utilizando a linguagem Go e abordarão diversas áreas críticas do sistema para garantir sua qualidade e robustez.

## Sistema de Broker MQTT

### Testes de Comunicação MQTT:
Verificar se o broker MQTT está recebendo e enviando mensagens corretamente.
Testar os casos de sucesso e falha para garantir a integridade da comunicação entre os clientes MQTT e o broker.
Garantir que as mensagens sejam entregues corretamente aos clientes inscritos nos tópicos relevantes.

### Testes de Segurança MQTT:
Testar a autenticação de clientes MQTT para garantir que apenas usuários autorizados tenham acesso ao sistema.
Verificar se as permissões de acesso aos tópicos estão sendo aplicadas corretamente para evitar acessos não autorizados.

Testar casos de ataques de negação de serviço (DoS) para avaliar a resiliência do sistema a esses tipos de ataques.
Servidor

## Servidor - Backend 

### Testes de Rotas HTTP:
Verificar se as rotas HTTP do servidor estão respondendo corretamente aos diferentes métodos HTTP (GET, POST, etc.).
Testar os casos de sucesso e falha para garantir o correto funcionamento das rotas protegidas e não protegidas.

### Testes de Segurança:
Testar os mecanismos de autenticação e autorização para garantir que apenas usuários autenticados tenham acesso aos recursos protegidos.
Verificar se o servidor está protegido contra ataques de injeção de código, XSS e CSRF.

## Considerações de Segurança

### Testes de Injeção de SQL:
Garantir que o servidor esteja protegido contra ataques de injeção de SQL, verificando se as consultas SQL são sanitizadas corretamente.

### Testes de XSS (Cross-Site Scripting):
Testar se o servidor está protegido contra ataques XSS, validando se os dados de entrada são corretamente sanitizados antes de serem renderizados no navegador.

### Testes de CSRF (Cross-Site Request Forgery):
Verificar se o servidor está protegido contra ataques CSRF, garantindo que todas as solicitações estejam acompanhadas de tokens de autenticação válidos.
Esses testes unitários ajudarão a garantir a qualidade e segurança do sistema EcoVigilância, permitindo uma implementação robusta e confiável. Certifique-se de incluir esses testes como parte do seu processo de desenvolvimento para garantir a integridade do sistema em todas as etapas.