# 📋 Instruções para Publicar no GitHub

## Pré-requisitos
- ✅ Git instalado e configurado localmente
- ✅ Conta GitHub criada

## Passos para Publicar

### 1. Criar um novo repositório no GitHub

1. Acesse [GitHub](https://github.com/new)
2. Preencha os campos:
   - **Repository name**: `desafio-fullstack-veritas`
   - **Description** (opcional): "Mini Kanban Fullstack - React + Go - Desafio Veritas"
   - **Visibility**: Public (recomendado para portfólio)
   - Deixe desmarcado "Initialize this repository with a README" (já temos um)
3. Clique em **Create repository**

### 2. Adicionar remote e fazer push

Copie e execute os comandos no terminal PowerShell:

```powershell
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas

# Adicionar o repositório remoto (substitua SEU_USUARIO)
git remote add origin https://github.com/SEU_USUARIO/desafio-fullstack-veritas.git

# Verificar que o remote foi adicionado
git remote -v

# Fazer push da branch master para origin
git branch -M main
git push -u origin main
```

### 3. Verificar no GitHub

Acesse `https://github.com/SEU_USUARIO/desafio-fullstack-veritas` para confirmar que todos os arquivos foram enviados.

---

## Estrutura do Repositório (o que foi enviado)

```
desafio-fullstack-veritas/
├── backend/
│   ├── go.mod                 # Módulo Go
│   ├── main.go               # Server setup
│   ├── handlers.go           # REST API endpoints
│   ├── models.go             # Task model
│   ├── storage.go            # Persistência JSON
│   └── handlers_test.go      # Testes unitários
├── frontend/
│   ├── package.json          # Dependencies (Vite + React)
│   ├── index.html            # HTML entry point
│   ├── src/
│   │   ├── main.jsx          # React entry
│   │   ├── App.jsx           # Main app (drag-drop)
│   │   ├── api.js            # API client
│   │   ├── styles.css        # Styling (DnD feedback)
│   │   └── components/
│   │       ├── TaskForm.jsx  # Form component
│   │       └── Toast.jsx     # Toast notifications
│   └── README.md             # Frontend setup
├── docs/
│   └── user-flow.svg         # Diagrama fluxo usuário
├── .gitignore                # Git ignore rules
├── README.md                 # Documentação principal
├── DRAG_DROP_TEST.md         # Guia de teste D&D
└── .git/                     # Git history (1 commit)
```

---

## Funcionalidades Implementadas

✅ **Backend (Go)**
- REST API em Go (std library, sem dependências)
- 5 endpoints: GET/POST /tasks, GET/PUT/DELETE /tasks/{id}
- Armazenamento em memória + persistência JSON (data.json)
- Validações: título obrigatório, status válido
- CORS configurado
- Testes unitários

✅ **Frontend (React)**
- Vite + React 18 (function components + hooks)
- 3 colunas fixas (A Fazer, Em Progresso, Concluídas)
- CRUD completo: criar, editar, mover, excluir
- **Drag-and-drop** entre colunas (HTML5 nativo)
- Toast notifications (sucesso/erro)
- Validação de entrada (título obrigatório)
- Loading & error states

✅ **Extras**
- Persistência em JSON (data.json)
- Testes unitários backend
- Documentação completa
- Diagrama user-flow (SVG)
- Drag-and-drop implementado

---

## Como Clonar e Rodar Localmente

Após a publicação, qualquer pessoa pode clonar e executar:

```powershell
# Clonar
git clone https://github.com/SEU_USUARIO/desafio-fullstack-veritas.git
cd desafio-fullstack-veritas

# Backend
cd backend
go run .

# Frontend (em outro terminal)
cd frontend
npm install
npm run dev
```

Acesse `http://localhost:5173` no navegador.

---

## Informações Adicionais

### Tecnologias Utilizadas
- **Backend**: Go 1.18+, stdlib (net/http)
- **Frontend**: React 18, Vite, CSS3 (HTML5 Drag & Drop)
- **Persistência**: JSON file (atomic writes)
- **Testing**: Go testing package

### Notas Importantes
- O repositório **não contém** `node_modules/`, `backend.exe`, `dist/` (listados em `.gitignore`)
- Primeiro push inclui 1 commit raiz com todo o projeto
- Para contribuições futuras, fazer commits pequenos com mensagens descritivas (ex: `feat: add filter by status`)

---

## Troubleshooting

### "Failed to connect to github.com"
- Verifique internet
- Tente usar SSH em vez de HTTPS:
  ```powershell
  git remote set-url origin git@github.com:SEU_USUARIO/desafio-fullstack-veritas.git
  ```

### "Permission denied (publickey)"
- Configure chave SSH: https://docs.github.com/en/authentication/connecting-to-github-with-ssh

### "Branch main not found"
- Pode estar em `master` ao invés de `main`. Verifique com:
  ```powershell
  git branch -a
  ```

---

**Pronto! Repositório preparado e instruções criadas.** 🚀
