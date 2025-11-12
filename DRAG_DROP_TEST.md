# Como Testar Drag-and-Drop

## Alterações Implementadas

O Kanban agora suporta **drag-and-drop nativo** para mover tarefas entre colunas.

### Arquivos Modificados

1. **`frontend/src/App.jsx`**
   - Adicionados 3 event handlers:
     - `handleDragStart(e, task)`: inicia o drag com os dados da tarefa
     - `handleDragOver(e)`: permite drop na coluna
     - `handleDrop(e, targetStatus)`: finaliza o movimento
   - Adicionados atributos `draggable` e event listeners nos cards de tarefa e colunas

2. **`frontend/src/styles.css`**
   - Adicionados estilos visuais:
     - `.task { cursor: move }` — indica que a tarefa pode ser arrastada
     - `.task:hover { background: #f0f0f0 }` — feedback ao hover
     - `.task[draggable="true"] { cursor: grab }` — cursor grab ao arrastar
     - `.column:hover { background: #f9fafb }` — destaque suave da coluna
     - `.column:has(.tasks:drag-over)` — destaca drop zone

3. **`README.md`**
   - Atualizado para mencionar drag-and-drop como feature principal

## Como Testar

### 1. Inicie o Backend e Frontend
```powershell
# Terminal 1: Backend (porta 8080)
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas\backend
go run .

# Terminal 2: Frontend (porta 5173)
cd C:\Users\PauloVictor\Desktop\desafio-fullstack-veritas\frontend
npm run dev
```

### 2. Abra o Navegador
Acesse: `http://localhost:5173/`

### 3. Teste Drag-and-Drop
1. **Crie uma tarefa**:
   - Digite título (ex: "Testar Drag")
   - Clique "Salvar"
   - Tarefa aparece em "A Fazer"

2. **Arraste a tarefa**:
   - Clique e segure a tarefa
   - Arraste para a coluna "Em Progresso"
   - Solte o mouse
   - ✅ Tarefa se move instantaneamente!

3. **Teste entre todas as colunas**:
   - "A Fazer" → "Em Progresso" → "Concluídas"
   - "Concluídas" → "A Fazer" (voltar)

### 4. Evidência de Sucesso
- ✅ Tarefa se move entre colunas ao arrastar
- ✅ Toast exibe "Tarefa atualizada" no canto superior direito
- ✅ Cursor muda para "grab" / "grabbing" ao arrastar
- ✅ Coluna fica destacada ao passar mouse
- ✅ Backend recebe PUT e persiste em `data.json`

## Comportamento Esperado

| Ação | Resultado |
|------|-----------|
| Arrastar tarefa para coluna diferente | Tarefa se move + toast sucesso |
| Arrastar tarefa para mesma coluna | Nenhuma ação (validação) |
| Soltar tarefa fora da coluna | Nenhuma ação |
| Criar tarefa → Arrastar | Funciona normalmente |
| Editar tarefa → Arrastar | Funciona normalmente |

## Notas Técnicas

- Implementação usa **HTML5 Drag and Drop API** (nativa, sem dependências)
- Compatível com todos os navegadores modernos (Chrome, Firefox, Safari, Edge)
- O movimento é imediato (otimistic update) sem aguardar resposta da API
- Se a API falhar, a UI exibe um toast de erro
- Arquivo `backend/data.json` persiste o novo status

## Possíveis Problemas & Soluções

| Problema | Solução |
|----------|---------|
| Tarefa não se move | Verifique se o backend está rodando (GET http://localhost:8080/tasks retorna []) |
| Erro "failed to update" | Verifique logs do backend e console do navegador (F12) |
| Cursor não muda | Limpe cache do navegador (Ctrl+Shift+Delete) |
| Toast não aparece | Console aberto (F12) mostra erros JavaScript |

---

**Pronto! Teste agora e reporte se houver problemas.**
