const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export async function fetchTasks() {
  const res = await fetch(`${API_BASE}/tasks`)
  if (!res.ok) throw new Error('failed to fetch tasks')
  return res.json()
}

export async function createTask(payload) {
  const res = await fetch(`${API_BASE}/tasks`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })
  if (!res.ok) {
    const txt = await res.text()
    throw new Error(txt || 'failed to create')
  }
  return res.json()
}

export async function updateTask(id, payload) {
  const res = await fetch(`${API_BASE}/tasks/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })
  if (!res.ok) {
    const txt = await res.text()
    throw new Error(txt || 'failed to update')
  }
  return res.json()
}

export async function deleteTask(id) {
  const res = await fetch(`${API_BASE}/tasks/${id}`, { method: 'DELETE' })
  if (!res.ok && res.status !== 204) {
    const txt = await res.text()
    throw new Error(txt || 'failed to delete')
  }
  return true
}
