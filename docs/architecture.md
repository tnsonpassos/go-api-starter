# Arquitetura

## Fluxo Principal

Request
↓
Handler
↓
Service
↓
Repository
↓
Database

## Camadas

### Handler

Responsável pela comunicação HTTP.

Funções:

* Receber requisições
* Validar entrada
* Retornar respostas

### Service

Responsável pelas regras de negócio.

Funções:

* Aplicar regras
* Orquestrar operações

### Repository

Responsável pela persistência.

Funções:

* Consultar dados
* Inserir dados
* Atualizar dados
* Remover dados

### Model

Representação das entidades do sistema.

## Estrutura Atual

internal/modules/items

* model.go
* handler.go
* service.go
* repository.go
* routes.go

pkg/response

* Padronização das respostas da API

## Banco de Dados

As alterações estruturais do banco são versionadas através de migrations.

Estrutura:

migrations/
├── 001_create_items.up.sql
└── 001_create_items.down.sql

Convenções:

- up = aplica alteração
- down = desfaz alteração

## Migration Runner

O projeto possui um comando separado para gerenciar alterações no banco.

Local:

cmd/migrate/main.go

Comandos:

go run cmd/migrate/main.go up
go run cmd/migrate/main.go down