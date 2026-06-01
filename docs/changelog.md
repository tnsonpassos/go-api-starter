# Changelog

Todas as mudanças relevantes deste projeto serão documentadas neste arquivo.

O formato utilizado segue o princípio:

* Added (Adicionado)
* Changed (Alterado)
* Fixed (Corrigido)
* Removed (Removido)

---

# v0.1.0 - 2026-05-31

Primeira versão funcional do Go API Starter.

## Added

### Estrutura Base

* Inicialização do módulo Go
* Estrutura de diretórios organizada
* Integração com Git e GitHub

### API

* Endpoint de Health Check
* CRUD completo de Items
* Rotas REST para:

  * Listagem
  * Consulta por ID
  * Criação
  * Atualização
  * Remoção

### Arquitetura

Implementação da arquitetura:

Request
→ Handler
→ Service
→ Repository
→ Data

Camadas criadas:

* Model
* Handler
* Service
* Repository
* Routes

### Configuração

* Configuração por arquivo .env
* Carregamento centralizado de configurações

### Responses

* Padronização das respostas JSON
* Estrutura única para sucesso e erro

### Documentação

Criação da documentação inicial:

* README
* Vision
* Architecture
* Roadmap
* Decisions
* Learning Journal
* Glossary
* Aprimorium

## Known Limitations

A versão atual ainda utiliza armazenamento em memória.

Os dados são perdidos ao reiniciar a aplicação.

## Next Version

v0.2.0 (em andamento)

Objetivos:

* Docker
* PostgreSQL
* Persistência real
* Repository conectado ao banco de dados

Added:
- PostgreSQL
- Repository com SQL
- Persistência real
- Integração Postgres.app
- Migration inicial