import type { Transaction } from '../types'

interface TransactionsTableProps {
  transactions: Transaction[]
  onDelete: (id: string) => Promise<void>
}

function formatCurrency(value: number, type: string) {
  const signed = type === 'expense' ? -value : value
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(signed)
}

function formatDate(date: string) {
  return new Intl.DateTimeFormat('en-US', {
    dateStyle: 'medium',
    timeStyle: 'short',
  }).format(new Date(date))
}

export function TransactionsTable({ transactions, onDelete }: TransactionsTableProps) {
  if (transactions.length === 0) {
    return <div className="empty-state">No transactions found for the current filters.</div>
  }

  return (
    <div className="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>Título</th>
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
              <td>{transaction.category}</td>
              <td>
                <span className={`badge ${transaction.type}`}>{transaction.type}</span>
              </td>
              <td>{formatCurrency(transaction.amount, transaction.type)}</td>
              <td>{formatDate(transaction.created_at)}</td>
              <td>
                <button className="danger-button" onClick={() => onDelete(transaction.id)}>
                  Deletar
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
