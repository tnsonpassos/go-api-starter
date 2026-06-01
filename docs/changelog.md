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

# v0.2.0 - 2026-05-31

## Added

- Integração com PostgreSQL
- Repository persistente para Items
- Migration inicial para tabela items
- Migration runner em Go
- Suporte a comando up
- Suporte a comando down
- Criação automática da tabela schema_migrations
- Makefile com atalhos de uso
- Validação do módulo Items
- Respostas de erro de validação padronizadas

## Changed

- Repository deixou de usar armazenamento em memória
- Handler passou a receber Service como dependência
- Routes passou a receber conexão com banco
- Fluxo principal passou a ser:
  Request → Handler → Service → Repository → PostgreSQL

## Known Limitations

- Ainda não há autenticação
- Ainda não há testes automatizados
- Ainda não há paginação
- Ainda não há Docker funcional neste ambiente

## Next Version

v0.3.0

Objetivos:

- Testes
- Logger
- Tratamento centralizado de erros
- Melhorias de qualidade