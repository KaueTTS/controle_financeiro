import type { Pagination } from '../../types/api';
import type { Transaction } from '../../types/transaction';
import { formatCurrency, formatDate, translateTransactionType } from '../../utils/formatters';

interface TransactionsTableProps {
  transactions: Transaction[];
  pagination: Pagination;
  isLoading: boolean;
  hasActiveFilters: boolean;
  onPageChange: (page: number) => void;
  onEdit: (transaction: Transaction) => void;
  onRequestDelete: (transaction: Transaction) => void;
}

export function TransactionsTable({
  transactions,
  pagination,
  isLoading,
  hasActiveFilters,
  onPageChange,
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
        <span className="total-count">{pagination.total} registro(s)</span>
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
        <>
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

                <time dateTime={transaction.createdAt} data-label="Data">
                  {formatDate(transaction.createdAt)}
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

          {pagination.pageCount > 1 && (
            <div className="pagination-actions">
              <button
                className="ghost-button small-button"
                type="button"
                disabled={pagination.page <= 1}
                onClick={() => onPageChange(pagination.page - 1)}
              >
                Anterior
              </button>

              <span>
                Página {pagination.page} de {pagination.pageCount}
              </span>

              <button
                className="ghost-button small-button"
                type="button"
                disabled={pagination.page >= pagination.pageCount}
                onClick={() => onPageChange(pagination.page + 1)}
              >
                Próxima
              </button>
            </div>
          )}
        </>
      )}
    </section>
  );
}