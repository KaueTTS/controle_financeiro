import type { Transaction } from '../../types/transaction';
import { formatCurrency, formatDate, translateTransactionType } from '../../utils/formatters';

interface TransactionsTableProps {
  transactions: Transaction[];
  isLoading: boolean;
  hasActiveFilters: boolean;
  onEdit: (transaction: Transaction) => void;
  onRequestDelete: (transaction: Transaction) => void;
}

export function TransactionsTable({
  transactions,
  isLoading,
  hasActiveFilters,
  onEdit,
  onRequestDelete,
}: TransactionsTableProps) {
  const emptyTitle = hasActiveFilters ? 'Nenhum registro encontrado' : 'Nenhum lançamento cadastrado';
  const emptyDescription = hasActiveFilters
    ? 'Ajuste os filtros para ampliar a busca.'
    : 'Cadastre um novo lançamento para começar.';

  return (
    <section className="panel table-panel">
      <div className="table-header">
        <div>
          <span className="eyebrow">Movimentações</span>
        </div>
        <span className="total-count">{transactions.length} registro(s)</span>
      </div>

      {isLoading ? (
        <div className="skeleton-list" aria-label="Carregando histórico">
          <span />
          <span />
          <span />
        </div>
      ) : transactions.length === 0 ? (
        <div className="empty-card">
          <strong>{emptyTitle}</strong>
          <p>{emptyDescription}</p>
        </div>
      ) : (
        <div className="transactions-list">
          <div className="transaction-row table-head" aria-hidden="true">
            <span>Lançamento</span>
            <span>Categoria</span>
            <span>Tipo</span>
            <span>Valor</span>
            <span>Data</span>
            <span>Ações</span>
          </div>

          {transactions.map((transaction) => (
            <article className="transaction-row" key={transaction.id}>
              <div className="transaction-main" data-label="Lançamento">
                <strong>{transaction.title}</strong>
                <small>{transaction.description || 'Sem descrição'}</small>
              </div>

              <span className="category-chip" data-label="Categoria">
                {transaction.category}
              </span>

              <span className={`pill ${transaction.type}`} data-label="Tipo">
                {translateTransactionType(transaction.type)}
              </span>

              <strong
                className={transaction.type === 'income' ? 'amount income-text' : 'amount expense-text'}
                data-label="Valor"
              >
                {transaction.type === 'expense' ? '-' : '+'} {formatCurrency(transaction.amount)}
              </strong>

              <time dateTime={transaction.created_at} data-label="Data">
                {formatDate(transaction.created_at)}
              </time>

              <div className="actions" data-label="Ações">
                <button className="ghost-button small-button" type="button" onClick={() => onEdit(transaction)}>
                  Editar
                </button>
                <button className="danger-button small-button" type="button" onClick={() => onRequestDelete(transaction)}>
                  Excluir
                </button>
              </div>
            </article>
          ))}
        </div>
      )}
    </section>
  );
}
