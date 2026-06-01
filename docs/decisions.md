# Registro de Decisões

## 2026-05-31

Decisão:
Usar Postgres.app temporariamente para a v0.2.0.

Motivo:
Docker Desktop e Colima apresentaram incompatibilidades no macOS 12.0.1 Intel.

Resultado:
Banco go_api_starter criado localmente com Postgres.app.

---

## 2026-05-31

Decisão:
Migrar para PostgreSQL via Docker.

Motivo:
Evitar problemas de instalação local.

---

## 2026-05-30

Decisão:
Separar arquitetura em Handler → Service → Repository.

Motivo:
Facilitar manutenção e reutilização.

---

## 2026-05-30

Decisão:
Utilizar "items" como módulo de exemplo.

Motivo:
Mais genérico que "products".