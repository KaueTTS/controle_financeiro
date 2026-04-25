import { useState } from 'react'
import type { TransactionType } from '../types'
import type { CreateTransactionPayload } from '../services/transactions'

interface TransactionFormProps {
  open: boolean
  loading: boolean
  onClose: () => void
  onSubmit: (payload: CreateTransactionPayload) => Promise<void>
}

const initialState: CreateTransactionPayload = {
  title: '',
  description: '',
  amount: 0,
  type: 'income',
  category: '',
}

export function TransactionForm({ open, loading, onClose, onSubmit }: TransactionFormProps) {
  const [form, setForm] = useState<CreateTransactionPayload>(initialState)

  if (!open) {
    return null
  }

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    await onSubmit(form)
    setForm(initialState)
  }

  function updateType(type: TransactionType) {
    setForm((current) => ({ ...current, type }))
  }

  return (
    <div className="modal-backdrop" role="presentation">
      <div className="modal">
        <div className="modal-header">
          <h2>Nova transação</h2>
          <button type="button" className="ghost-button" onClick={onClose}>
            Fechar
          </button>
        </div>

        <form onSubmit={handleSubmit} className="form-grid">
          <input
            required
            minLength={3}
            value={form.title}
            onChange={(event) => setForm((current) => ({ ...current, title: event.target.value }))}
            placeholder="Título"
          />

          <input
            value={form.description}
            onChange={(event) => setForm((current) => ({ ...current, description: event.target.value }))}
            placeholder="Descrição"
          />

          <input
            required
            min={0.01}
            step="0.01"
            type="number"
            value={form.amount || ''}
            onChange={(event) =>
              setForm((current) => ({
                ...current,
                amount: Number(event.target.value),
              }))
            }
            placeholder="Valor"
          />

          <div className="type-selector">
            <button
              type="button"
              className={form.type === 'income' ? 'active' : ''}
              onClick={() => updateType('income')}
            >
              Renda
            </button>
            <button
              type="button"
              className={form.type === 'expense' ? 'active' : ''}
              onClick={() => updateType('expense')}
            >
              Despesa
            </button>
          </div>

          <input
            required
            minLength={2}
            value={form.category}
            onChange={(event) => setForm((current) => ({ ...current, category: event.target.value }))}
            placeholder="Categoria"
          />

          <button type="submit" disabled={loading} className="primary-button">
            {loading ? 'Salvando...' : 'Salvar transação'}
          </button>
        </form>
      </div>
    </div>
  )
}
