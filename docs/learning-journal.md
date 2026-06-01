# Diário de Aprendizado

## 2026-05-31

Aprendizado:
O Repository agora abstrai a persistência.

Antes:
Slice em memória

Agora:
PostgreSQL

Benefício:
A camada superior não precisa saber onde os dados são armazenados.

---

Aprendizado:
Migration runner é um comando separado da API.

Função:
Executar alterações estruturais no banco de dados.

Fluxo:
Ler arquivos SQL
→ verificar schema_migrations
→ executar migrations pendentes
→ registrar versão aplicada

Observação:
A tabela schema_migrations funciona como histórico do banco.

## 2026-05-30

Aprendizado:
go mod init cria o módulo do projeto.

Observação:
Equivale ao package.json do Node.

---

Aprendizado:
Repository Pattern.

Função:
Separar acesso a dados das regras de negócio.

---

Aprendizado:
Configuração por .env.

Benefício:
Permite ambientes diferentes sem alterar código.