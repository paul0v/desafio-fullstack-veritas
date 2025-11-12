import React, { useEffect, useState } from 'react'
import { fetchTasks, createTask, updateTask, deleteTask } from './api'
import TaskForm from './components/TaskForm'
import Toast from './components/Toast'

const STATUS_LABELS = {
  todo: 'A Fazer',
  in_progress: 'Em Progresso',
  done: 'Concluídas'
}

export default function App() {
  const [tasks, setTasks] = useState([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [editing, setEditing] = useState(null)
  const [toasts, setToasts] = useState([])

  function pushToast(message, type = 'info', duration = 3000) {
    const id = Date.now() + Math.floor(Math.random() * 1000)
    setToasts(prev => [...prev, { id, message, type, duration }])
  }

  function removeToast(id) {
    setToasts(prev => prev.filter(t => t.id !== id))
  }

  async function load() {
    setLoading(true)
    setError(null)
    try {
      const data = await fetchTasks()
      setTasks(data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { load() }, [])

  async function handleCreate(payload) {
    setLoading(true)
    try {
      const t = await createTask(payload)
      setTasks(prev => [...prev, t])
      pushToast('Tarefa criada', 'success')
    } catch (err) {
      setError(err.message)
      pushToast(err.message, 'error')
    } finally { setLoading(false) }
  }

  async function handleUpdate(id, payload) {
    setLoading(true)
    try {
      const t = await updateTask(id, payload)
      setTasks(prev => prev.map(p => (p.id === id ? t : p)))
      setEditing(null)
      pushToast('Tarefa atualizada', 'success')
    } catch (err) {
      setError(err.message)
      pushToast(err.message, 'error')
    } finally { setLoading(false) }
  }

  async function handleDelete(id) {
    if (!confirm('Excluir tarefa?')) return
    setLoading(true)
    try {
      await deleteTask(id)
      setTasks(prev => prev.filter(p => p.id !== id))
      pushToast('Tarefa excluída', 'success')
    } catch (err) {
      setError(err.message)
      pushToast(err.message, 'error')
    } finally { setLoading(false) }
  }

  function moveTask(task, targetStatus) {
    if (task.status === targetStatus) return
    handleUpdate(task.id, { ...task, status: targetStatus })
  }

  // Drag-and-drop handlers
  function handleDragStart(e, task) {
    e.dataTransfer.effectAllowed = 'move'
    e.dataTransfer.setData('taskId', task.id)
  }

  function handleDragOver(e) {
    e.preventDefault()
    e.dataTransfer.dropEffect = 'move'
  }

  function handleDrop(e, targetStatus) {
    e.preventDefault()
    const taskId = parseInt(e.dataTransfer.getData('taskId'), 10)
    const task = tasks.find(t => t.id === taskId)
    if (task && task.status !== targetStatus) {
      moveTask(task, targetStatus)
    }
  }

  const byStatus = (s) => tasks.filter(t => t.status === s)

  return (
    <div className="app">
      <header>
        <h1>Mini Kanban</h1>
      </header>

      <section className="controls">
        <h2>Nova tarefa</h2>
        <TaskForm onSubmit={handleCreate} loading={loading} />
      </section>

      {loading && <div className="msg">Carregando...</div>}
      {error && <div className="msg error">{error}</div>}

      <main className="board">
        {Object.keys(STATUS_LABELS).map(status => (
          <div key={status} className="column" onDragOver={handleDragOver} onDrop={(e) => handleDrop(e, status)}>
            <h3>{STATUS_LABELS[status]}</h3>
            <div className="tasks">
              {byStatus(status).map(t => (
                <div key={t.id} className="task" draggable onDragStart={(e) => handleDragStart(e, t)}>
                  <div className="task-header">
                    <strong>{t.title}</strong>
                    <div className="task-actions">
                      <button onClick={() => setEditing(t)} disabled={loading}>Editar</button>
                      <button onClick={() => handleDelete(t.id)} disabled={loading}>Excluir</button>
                    </div>
                  </div>
                  {t.description && <p className="desc">{t.description}</p>}
                  <div className="move-controls">
                    {status !== 'todo' && (
                      <button onClick={() => moveTask(t, 'todo')} disabled={loading}>Mover para A Fazer</button>
                    )}
                    {status !== 'in_progress' && (
                      <button onClick={() => moveTask(t, 'in_progress')} disabled={loading}>Mover para Em Progresso</button>
                    )}
                    {status !== 'done' && (
                      <button onClick={() => moveTask(t, 'done')} disabled={loading}>Mover para Concluídas</button>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </div>
        ))}
      </main>

      {editing && (
        <div className="modal">
          <div className="modal-content">
            <h3>Editar tarefa</h3>
            <TaskForm
              initial={editing}
              loading={loading}
              onSubmit={(payload) => handleUpdate(editing.id, { ...editing, ...payload })}
              onCancel={() => setEditing(null)}
            />
          </div>
        </div>
      )}

      <div className="toasts-wrapper">
        {toasts.map(t => (
          <Toast key={t.id} id={t.id} message={t.message} type={t.type} duration={t.duration} onClose={removeToast} />
        ))}
      </div>
    </div>
  )
}
