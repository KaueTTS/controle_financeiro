import { useState, type FormEvent } from 'react';
import type { Transaction, TransactionPayload, TransactionType } from '../../types/transaction';

interface TransactionFormProps {
  transactionToEdit: Transaction | null;
  onSubmit: (payload: TransactionPayload, id?: number) => Promise<void>;
  onCancelEdit: () => void;
}

interface TransactionFormState {
  title: string;
  description: string;
  amount: string;
  category: string;
  type: TransactionType;
}

const EMPTY_FORM: TransactionFormState = {
  title: '',
  description: '',
  amount: '',
  category: '',
  type: 'expense',
};

function getInitialForm(transaction: Transaction | null): TransactionFormState {
  if (!transaction) {
    return EMPTY_FORM;
  }

  return {
    title: transaction.title,
    description: transaction.description ?? '',
    amount: String(transaction.amount),
    category: transaction.category,
    type: transaction.type,
  };
}

export function TransactionForm({ transactionToEdit, onSubmit, onCancelEdit }: TransactionFormProps) {
  const [form, setForm] = useState(() => getInitialForm(transactionToEdit));
  const [isSubmitting, setIsSubmitting] = useState(false);

  const isEditing = Boolean(transactionToEdit);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const amount = Number(form.amount);

    if (!form.title.trim() || !form.category.trim() || amount <= 0) {
      return;
    }

    const payload: TransactionPayload = {
      title: form.title.trim(),
      description: form.description.trim() || null,
      amount,
      category: form.category.trim(),
      type: form.type,
    };

    try {
      setIsSubmitting(true);
      await onSubmit(payload, transactionToEdit?.id);
      setForm(EMPTY_FORM);
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <form className="transaction-form" onSubmit={handleSubmit}>
      <div className="type-toggle" aria-label="Tipo do lançamento">
        <button
          type="button"
          className={form.type === 'expense' ? 'active expense-option' : ''}
          onClick={() => setForm({ ...form, type: 'expense' })}
        >
          Despesa
        </button>
        <button
          type="button"
          className={form.type === 'income' ? 'active income-option' : ''}
          onClick={() => setForm({ ...form, type: 'income' })}
        >
          Receita
        </button>
      </div>

      <div className="form-grid two-columns">
        <div className="field field-span-2">
          <label htmlFor="title">Título</label>
          <input
            id="title"
            type="text"
            placeholder="Ex: Salário, mercado, aluguel"
            value={form.title}
            onChange={(event) => setForm({ ...form, title: event.target.value })}
            required
          />
        </div>

        <div className="field">
          <label htmlFor="amount">Valor</label>
          <input
            id="amount"
            type="number"
            min="0.01"
            step="0.01"
            placeholder="0,00"
            value={form.amount}
            onChange={(event) => setForm({ ...form, amount: event.target.value })}
            required
          />
        </div>

        <div className="field">
          <label htmlFor="category-form">Categoria</label>
          <input
            id="category-form"
            type="text"
            placeholder="Ex: Casa"
            value={form.category}
            onChange={(event) => setForm({ ...form, category: event.target.value })}
            required
          />
        </div>

        <div className="field field-span-2">
          <label htmlFor="description">Descrição</label>
          <textarea
            id="description"
            placeholder="Opcional"
            value={form.description}
            onChange={(event) => setForm({ ...form, description: event.target.value })}
            rows={3}
          />
        </div>
      </div>

      <div className="modal-actions">
        <button className="ghost-button" type="button" onClick={onCancelEdit}>
          Cancelar
        </button>
        <button className="primary-button" type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Salvando...' : isEditing ? 'Salvar alterações' : 'Adicionar'}
        </button>
      </div>
    </form>
  );
}
