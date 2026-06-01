# Roadmap

## Visão Geral

O Go API Starter é um template reutilizável para desenvolvimento de APIs REST em Go.

Objetivos:

* Aprender Go de forma prática
* Criar uma arquitetura reutilizável
* Reduzir o tempo de criação de novos projetos
* Evoluir para um starter pronto para produção

---

# Versão Atual

## v0.1.0 — Starter em Memória ✅

Status: Concluído

### Funcionalidades

* [x] Estrutura inicial do projeto
* [x] CRUD Items
* [x] Handler Layer
* [x] Service Layer
* [x] Repository Layer
* [x] Response Package
* [x] Configuração via .env
* [x] GitHub
* [x] Documentação inicial

### Resultado

Primeira versão funcional da arquitetura.

---

# v0.2.0 — Persistência Real

Status: Concluído

### Ambiente

* [ ] Docker Desktop
* [ ] Docker Compose
* [x] PostgreSQL Container

### Backend

* [x] Conexão PostgreSQL
* [x] Health Check do Banco
* [x] Repository PostgreSQL
* [x] Remoção do armazenamento em memória

### Dados

* [x] Tabela Items
* [x] CRUD persistente

### Documentação

* [ ] Guia Docker
* [ ] Guia PostgreSQL

### Resultado Esperado

Dados persistidos em banco de dados.

---

# v0.3.0 — Migrations e Qualidade

Status: Planejado

### Banco

* [x] Sistema de Migrations
* [x] Seed de dados

### Qualidade

* [ ] Logger
* [ ] Tratamento centralizado de erros
* [ ] Testes unitários

### Resultado Esperado

Base sólida para crescimento do projeto.

---

# v0.4.0 — Autenticação

Status: Planejado

### Usuários

* [ ] Model User
* [ ] CRUD User

### Segurança

* [ ] JWT
* [ ] Login
* [ ] Registro
* [ ] Middleware de autenticação

### Resultado Esperado

API protegida por autenticação.

---

# v0.5.0 — Starter Reutilizável

Status: Planejado

### Automação

* [ ] Gerador de módulos
* [ ] Gerador de CRUD

### Estrutura

* [ ] Interfaces para repositories
* [ ] Injeção de dependência

### Resultado Esperado

Capacidade de criar novos módulos rapidamente.

---

# v1.0.0 — Starter Production Ready

Status: Futuro

### Infraestrutura

* [ ] Docker Compose completo
* [ ] Ambiente Dev
* [ ] Ambiente Produção

### Segurança

* [ ] Rate Limiting
* [ ] CORS
* [ ] Auditoria

### Observabilidade

* [ ] Logs estruturados
* [ ] Métricas

### Resultado Esperado

Starter pronto para projetos reais.

---

# Backlog

Ideias para avaliação futura:

* Swagger/OpenAPI
* Refresh Tokens
* RBAC (Roles e Permissões)
* Upload de Arquivos
* Cache Redis
* Filas
* WebSockets
* CLI de scaffolding
* Template para Microsserviços
* Template SaaS
* Template Admin Panel

```
```