import React, { useState, useEffect } from 'react'

export default function TaskForm({ onSubmit, initial, onCancel, loading }) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [touched, setTouched] = useState(false)

  useEffect(() => {
    if (initial) {
      setTitle(initial.title || '')
      setDescription(initial.description || '')
    }
  }, [initial])

  function submit(e) {
    e.preventDefault()
    setTouched(true)
    if (!title.trim()) return
    onSubmit({ title: title.trim(), description: description.trim(), status: initial?.status })
    // only clear when creating (no initial provided)
    if (!initial) {
      setTitle('')
      setDescription('')
      setTouched(false)
    }
  }

  return (
    <form className="task-form" onSubmit={submit}>
      <input
        placeholder="Título (obrigatório)"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        onBlur={() => setTouched(true)}
      />
      {touched && !title.trim() && <div className="input-error">Título é obrigatório</div>}
      <textarea
        placeholder="Descrição (opcional)"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />
      <div className="form-actions">
        <button type="submit" disabled={loading || !title.trim()}>{loading ? 'Salvando...' : 'Salvar'}</button>
        {onCancel && (
          <button type="button" className="btn-secondary" onClick={onCancel} disabled={loading}>
            Cancelar
          </button>
        )}
      </div>
    </form>
  )
}
