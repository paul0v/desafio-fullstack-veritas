# Frontend (Mini Kanban)

This is a Vite + React frontend for the Veritas mini-kanban challenge.

How to run

1. Install dependencies

```powershell
cd frontend
npm install
```

2. Start dev server

```powershell
npm run dev
```

The app expects the backend at `http://localhost:8080` by default. You can override by setting `VITE_API_URL`.

Notes

- Minimal UI focused on functionality: create, edit, move, delete tasks.
- Uses simple fetch wrappers in `src/api.js` to call the backend.
