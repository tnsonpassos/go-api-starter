# Go API Starter — Frontend

Frontend criado com Vue 3, Vite e Bootstrap para consumir a API Go API Starter.

## Stack

- Vue 3
- Vite
- Bootstrap
- Fetch API

## Objetivo

Criar uma interface visual para testar e consumir o CRUD de Items da API Go.

## Funcionalidades

- [x] Listar Items
- [x] Criar Item
- [x] Editar Item
- [x] Excluir Item
- [x] Mensagens de sucesso
- [x] Mensagens de erro
- [x] Estado de edição
- [x] Estado de lista vazia

## Como rodar

### Na raiz do projeto, rode o backend:

```bash
make run
```

### Em outro terminal, rode o frontend:

```bash
cd web
npm run dev
```

### Frontend:

http://localhost:5173

### Backend:

http://localhost:8080

## Estrutura

```txt
web/
├── public/
├── src/
│   ├── components/
│   ├── pages/
│   │   └── ItemsPage.vue
│   ├── services/
│   │   └── api.js
│   ├── App.vue
│   └── main.js
├── package.json
├── vite.config.js
└── README.md
```

## Integração com Backend

O frontend consome os endpoints:

GET    /items

POST   /items

PUT    /items/:id

DELETE /items/:id

## Observações

A API precisa estar rodando antes do frontend consumir os dados.

O backend precisa liberar CORS para:

http://localhost:5173