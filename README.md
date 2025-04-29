# API em Go para Estudos ✨

Esta API foi desenvolvida utilizando o framework **Gin** e conta com um banco de dados **PostgreSQL** executado dentro de um contêiner **Docker**. O objetivo principal do projeto é o aprendizado e prática de conceitos de desenvolvimento backend com Go.

## ⚡ Tecnologias Utilizadas

- **Go**: Linguagem principal da API.
- **Gin**: Framework para simplificar o desenvolvimento web.
- **Docker**: Utilizado para rodar uma imagem do PostgreSQL.
- **PostgreSQL**: Banco de dados relacional.
- **Swagger**: Documentação da API.
- **JWT**: Autenticação de usuários com tokens seguros. 🔐

## 🌟 Objetivo 😊

Este projeto foi criado com o propósito de estudo e aprendizado. Não deve ser utilizado em produção. Ele serve como base para explorar conceitos como:

- Criação de rotas e handlers com Gin.
- Conexão com o banco de dados PostgreSQL.
- Utilização do Docker para gerenciar serviços de forma isolada.
- Implementação de autenticação utilizando JWT. 🔑

## 📂 Estrutura do Projeto

A estrutura do projeto está organizada da seguinte forma:

```
API-MERCEARIA-GO/
├── controller/         # Contém os handlers das rotas
├── db/                # Configurações e migrações do banco de dados
├── docs/              # Documentação gerada pelo Swagger
├── model/            # Modelos que representam as entidades
├── repository/       # Interações com o banco de dados
├── usecase/          # Lógica de negócio e casos de uso
├── docker-compose.yaml  # Arquivo para orquestração com Docker
├── Dockerfile         # Arquivo para build da imagem Docker
├── go.mod            # Gerenciamento de dependências do Go
├── go.sum            # Checksum das dependências
├── main.go          # Ponto de entrada da aplicação
└── README.md         # Documentação do projeto
```

## 🚀 Contribuições

Como este projeto é voltado para aprendizado, você é bem-vindo(a) para contribuir, sugerir melhorias ou relatar problemas. Sinta-se à vontade para abrir issues ou enviar pull requests. 📢

