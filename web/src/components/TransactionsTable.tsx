import type { Transaction } from '../types'

interface TransactionsTableProps {
  transactions: Transaction[]
  onDelete: (id: string) => Promise<void>
}

function formatCurrency(value: number, type: string) {
  const signed = type === 'expense' ? -value : value
  return new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  }).format(signed)
}

function formatDate(date: string) {
  return new Intl.DateTimeFormat('pt-BR', {
    dateStyle: 'medium',
    timeStyle: 'short',
  }).format(new Date(date))
}

function translateType(type: string) {
  const map: Record<string, string> = {
    income: 'Receita',
    expense: 'Despesa',
  }

  return map[type] ?? type
}

export function TransactionsTable({ transactions, onDelete }: TransactionsTableProps) {
  if (transactions.length === 0) {
    return <div className="empty-state">Sem transações para exibir.</div>
  }

  return (
    <div className="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>Título</th>
            <th>Descrição</th>
            <th>Categoria</th>
            <th>Tipo</th>
            <th>Valor</th>
            <th>Criado em</th>
            <th />
          </tr>
        </thead>
        <tbody>
          {transactions.map((transaction) => (
            <tr key={transaction.id}>
              <td>{transaction.title}</td>
              <td>{transaction.description}</td>
              <td>{transaction.category}</td>
              <td>
                <span className={`badge ${transaction.type}`}>{translateType(transaction.type)}</span>
              </td>
              <td>{formatCurrency(transaction.amount, transaction.type)}</td>
              <td>{formatDate(transaction.created_at)}</td>
              <td>
                <div className="actions">
                  <button className="danger-button" onClick={() => onDelete(transaction.id)}>
                    Deletar
                  </button>
                  <button
                    className="edit-button"
                    onClick={() => alert('Funcionalidade de edição ainda não implementada')}
                  >
                    Editar
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
