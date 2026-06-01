# Go API Starter

## Visão Geral

O Go API Starter é um projeto de estudo e desenvolvimento criado com dois objetivos principais:

1. Aprender Go de forma prática através da construção de uma API real.
2. Criar um template reutilizável para acelerar futuros projetos.

O projeto evoluirá gradualmente para incluir:

* CRUD reutilizável
* Configuração por ambiente
* PostgreSQL
* Docker
* Migrations
* Autenticação JWT
* Controle de permissões
* Geração automática de módulos

## Comandos

### Rodar API

```bash
go run cmd/api/main.go
```

### Migrations

- Rodar todas as migrations pendentes

```bash
go run cmd/migrate/main.go up
```

- Reverter a última migration

```bash
go run cmd/migrate/main.go down
```

### Testes

- Testar health check

```bash
curl http://localhost:8080/health
```

- Listar items

```bash
curl http://localhost:8080/items
```

## Makefile

O projeto possui atalhos para comandos comuns.

```bash
make run
```

```bash
make migrate-up
```

```bash
make migrate-down
```

```bash
make test-health
```

```bash
make test-items
```

## Estado Atual

### Concluído

* Estrutura inicial do projeto
* CRUD Items
* Arquitetura Handler → Service → Repository
* Configuração por .env
* Padronização de respostas JSON
* Repositório GitHub

### Próximos Passos

* Docker
* PostgreSQL
* Persistência real
* Migrations

## Filosofia

Este projeto segue os princípios do Aprimorium:

Função → Forma → Fluxo

Todo componente deve possuir:

* Uma função clara
* Uma forma simples
* Um fluxo compreensível