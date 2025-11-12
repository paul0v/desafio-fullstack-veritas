# Desafio Fullstack - Mini Kanban (Veritas)

Repositório com um pequeno Kanban (React frontend + Go backend) desenvolvido para o desafio técnico da Veritas Consultoria.

Sumário
- Visão geral
- Estrutura do repositório
- Como rodar (backend / frontend)
- API (endpoints principais)
- Decisões técnicas
- Limitações conhecidas & melhorias futuras
- Documentação (diagramas)

## Visão geral

Aplicação mínima (MVP) que permite criar, editar, mover entre colunas e excluir tarefas em um quadro com três colunas fixas: A Fazer, Em Progresso e Concluídas. Os dados são servidos por uma API REST em Go. Suporta **drag-and-drop** entre colunas para mover tarefas de forma intuitiva.

## Estrutura do repositório

```
desafio-fullstack-veritas/
├─ backend/        # servidor Go (main.go, handlers.go, models.go)
├─ frontend/       # app React (Vite)
└─ docs/           # diagramas (user-flow.png etc.)
```

## Como rodar

Requisitos:
- Go (>= 1.18)
- Node.js + npm (para frontend)

1) Backend (Go)

Abra um terminal PowerShell e execute:

```powershell
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas\backend
go build
.\backend.exe
```

Ou para rodar sem compilar:

```powershell
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas\backend
go run .
```

O servidor inicia em http://localhost:8080 e expõe os endpoints REST descritos abaixo. CORS está configurado para permitir requests do frontend durante desenvolvimento.

2) Frontend (Vite + React)

Instale dependências e inicie o dev server:

```powershell
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas\frontend
npm install
npm run dev
```

O Vite normalmente abre em http://localhost:5173. O frontend faz requisições por padrão para `http://localhost:8080`; se seu backend estiver em outro endereço, exporte a variável antes de iniciar:

```powershell
$env:VITE_API_URL='http://seu-backend:8080'; npm run dev
```

## API (resumo)

Base URL: `http://localhost:8080`

- GET /tasks
  - Retorna lista de tarefas (200)
- POST /tasks
  - Cria nova tarefa (201). Body JSON: { title: string (obrigatório), description?: string, status?: "todo"|"in_progress"|"done" }
- GET /tasks/{id}
  - Retorna tarefa por id (200 / 404)
- PUT /tasks/{id}
  - Atualiza tarefa (200 / 404). Mesmas validações do POST (título obrigatório, status válido)
- DELETE /tasks/{id}
  - Remove a tarefa (204 / 404)

Validações principais: título obrigatório; status deve ser um dos valores permitidos. Em caso de erro a API retorna um status HTTP apropriado e mensagem simples no corpo.

## Decisões técnicas

- Backend em Go (std library net/http): leve, sem dependências externas para simplicidade. Armazenamento atual: memória com persistência em JSON.
- CORS permissivo durante desenvolvimento para facilitar integração com Vite.
- Frontend com Vite + React (function components + hooks) para desenvolvimento rápido e bundle leve.
- **Drag-and-drop**: implementado com HTML5 Drag and Drop API nativa (sem dependências). Clique e arraste uma tarefa entre colunas para movê-la instantaneamente.
- Organização: separação clara entre `backend/` e `frontend/`, e `docs/` para diagramas.

## Limitações conhecidas e melhorias futuras

- Autenticação/Autorização: não implementada (fora do escopo do MVP).
- Testes E2E: falta suíte de testes de integração (backend tem testes unitários).
- UI/UX avançadas: podemos adicionar melhores transições e animações.
- Docker: não incluído, mas podemos adicionar Dockerfile(s) para backend e frontend.

## Documentação / Diagramas

- `docs/user-flow.png` — diagrama obrigatório (ainda não adicionado). Coloque o arquivo em `docs/` com esse nome.
- `docs/data-flow.png` — opcional para mostrar comunicação entre frontend, API e armazenamento.

Se quiser, eu crio um esboço do `user-flow.png` e adiciono em `docs/`.

## Como contribuir / preparar para envio

- Faça commits pequenos e explicativos, por exemplo: `backend: add tasks handlers`, `frontend: add TaskForm component`.
- Antes de subir para o GitHub verifique:
  - Backend compila (go build)
  - Frontend instala dependências (npm install)
  - Atualize/adicione `docs/user-flow.png`

---

Se quiser, eu implemento agora:
- Persistência em arquivo JSON no backend (rápido) — ou
- Gerar um esboço `docs/user-flow.png` e adicioná-lo ao repositório — ou
- Adicionar testes simples ao backend.

Diga qual desses prefere que eu execute a seguir.
