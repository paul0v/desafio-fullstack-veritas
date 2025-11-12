import React, { useEffect } from 'react'

export default function Toast({ id, message, type = 'info', onClose, duration = 3000 }) {
  useEffect(() => {
    const t = setTimeout(() => onClose(id), duration)
    return () => clearTimeout(t)
  }, [id, duration, onClose])

  return (
    <div className={`toast toast-${type}`} role="status">
      {message}
    </div>
  )
}
