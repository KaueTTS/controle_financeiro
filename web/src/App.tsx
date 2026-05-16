import { useEffect, useMemo, useState } from 'react';
import { ConfirmDialog } from './components/common/ConfirmDialog';
import { AppHeader } from './components/layout/AppHeader';
import { SummaryCards } from './components/summary/SummaryCards';
import { TransactionFilters } from './components/transactions/TransactionFilters';
import { TransactionForm } from './components/transactions/TransactionForm';
import { TransactionsTable } from './components/transactions/TransactionsTable';
import { THEME_STORAGE_KEY } from './constants/dashboard';
import { useTransactions } from './hooks/useTransactions';
import type { Theme } from './types/theme';
import type { Transaction } from './types/transaction';
import { formatCurrency } from './utils/formatters';

function App() {
  const [transactionToEdit, setTransactionToEdit] = useState<Transaction | null>(null);
  const [transactionToDelete, setTransactionToDelete] = useState<Transaction | null>(null);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [theme, setTheme] = useState<Theme>(() => {
    const savedTheme = localStorage.getItem(THEME_STORAGE_KEY);
    return savedTheme === 'dark' ? 'dark' : 'light';
  });

  const {
    transactions,
    summary,
    filters,
    categories,
    isLoading,
    isDeleting,
    error,
    hasActiveFilters,
    setFilters,
    clearFilters,
    saveTransaction,
    removeTransaction,
  } = useTransactions();

  const hasPositiveBalance = summary.balance >= 0;

  const financialStatus = useMemo(() => {
    const hasSummaryData = summary.income > 0 || summary.expense > 0;

    if (!hasSummaryData) {
      return hasActiveFilters
        ? 'Nenhum resultado encontrado para os filtros aplicados.'
        : 'Comece cadastrando seu primeiro lançamento.';
    }

    if (summary.balance > 0) {
      return 'Você tem dinheiro disponível no período selecionado.';
    }

    if (summary.balance < 0) {
      return 'As despesas superaram as receitas neste período.';
    }

    return 'Suas receitas e despesas estão equilibradas neste período.';
  }, [hasActiveFilters, summary.balance, summary.expense, summary.income]);

  useEffect(() => {
    document.documentElement.dataset.theme = theme;
    localStorage.setItem(THEME_STORAGE_KEY, theme);
  }, [theme]);

  function toggleTheme() {
    setTheme((currentTheme) => (currentTheme === 'light' ? 'dark' : 'light'));
  }

  function openCreateModal() {
    setTransactionToEdit(null);
    setIsFormOpen(true);
  }

  function openEditModal(transaction: Transaction) {
    setTransactionToEdit(transaction);
    setIsFormOpen(true);
  }

  function closeFormModal() {
    setTransactionToEdit(null);
    setIsFormOpen(false);
  }

  function closeDeleteDialog() {
    if (!isDeleting) {
      setTransactionToDelete(null);
    }
  }

  async function handleSave(...params: Parameters<typeof saveTransaction>) {
    await saveTransaction(...params);
    closeFormModal();
  }

  async function handleConfirmDelete() {
    if (!transactionToDelete) {
      return;
    }

    try {
      await removeTransaction(transactionToDelete.id);
      setTransactionToDelete(null);
    } catch {
      // O hook já exibe a mensagem de erro no alerta principal.
    }
  }

  return (
    <main className="app-shell">
      <AppHeader theme={theme} onToggleTheme={toggleTheme} onCreateTransaction={openCreateModal} />

      {error && <div className="alert">{error}</div>}

      <section className="hero-card">
        <div className="hero-copy">
          <span className="eyebrow">Saldo atual</span>

          <strong className={hasPositiveBalance ? 'income-text' : 'expense-text'}>
            {formatCurrency(summary.balance)}
          </strong>

          <div className={`balance-status ${hasPositiveBalance ? 'positive' : 'negative'}`}>
            <span>{hasPositiveBalance ? '✓' : '!'}</span>
            {hasPositiveBalance ? 'Seu saldo está positivo' : 'Seu saldo está negativo'}
          </div>

          <p>{financialStatus}</p>
        </div>

        <SummaryCards summary={summary} />
      </section>

      <section className="content-card">
        <TransactionFilters
          filters={filters}
          categories={categories}
          hasActiveFilters={hasActiveFilters}
          onChange={setFilters}
          onClear={clearFilters}
        />

        <TransactionsTable
          transactions={transactions}
          isLoading={isLoading}
          hasActiveFilters={hasActiveFilters}
          onEdit={openEditModal}
          onRequestDelete={setTransactionToDelete}
        />
      </section>

      {isFormOpen && (
        <div className="modal-backdrop" role="presentation" onMouseDown={closeFormModal}>
          <div
            className="modal-card"
            role="dialog"
            aria-modal="true"
            aria-labelledby="transaction-modal-title"
            onMouseDown={(event) => event.stopPropagation()}
          >
            <div className="modal-header">
              <div>
                <span className="eyebrow">{transactionToEdit ? 'Edição' : 'Cadastro'}</span>
                <h2 id="transaction-modal-title">
                  {transactionToEdit ? 'Editar lançamento' : 'Novo lançamento'}
                </h2>
              </div>

              <button className="icon-button" type="button" onClick={closeFormModal} aria-label="Fechar modal">
                ×
              </button>
            </div>

            <TransactionForm
              transactionToEdit={transactionToEdit}
              onSubmit={handleSave}
              onCancelEdit={closeFormModal}
            />
          </div>
        </div>
      )}

      <ConfirmDialog
        isOpen={Boolean(transactionToDelete)}
        title="Excluir lançamento?"
        description={`Essa ação removerá "${transactionToDelete?.title ?? 'este lançamento'}" do histórico.`}
        confirmLabel="Excluir"
        isConfirming={isDeleting}
        onConfirm={handleConfirmDelete}
        onCancel={closeDeleteDialog}
      />
    </main>
  );
}

export default App;
