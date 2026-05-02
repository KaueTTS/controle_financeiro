import { useCallback, useEffect, useState } from 'react'
import { FilterBar } from '../components/FilterBar'
import { SummaryCards } from '../components/SummaryCards'
import { TransactionForm } from '../components/TransactionForm'
import { TransactionsTable } from '../components/TransactionsTable'
import {
  updateTransaction,
  createTransaction,
  deleteTransaction,
  fetchSummary,
  fetchTransactions,
  type CreateTransactionPayload,
} from '../services/transactions'
import type { Filters, Summary, Transaction } from '../types'

const initialFilters: Filters = {
  search: '',
  type: '',
  category: '',
}

const emptySummary: Summary = {
  income: 0,
  expense: 0,
  balance: 0,
}

export function HomePage() {
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [summary, setSummary] = useState<Summary>(emptySummary)
  const [filters, setFilters] = useState<Filters>(initialFilters)
  const [isLoading, setIsLoading] = useState(true)
  const [isSaving, setIsSaving] = useState(false)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [selectedTransaction, setSelectedTransaction] = useState<Transaction | null>(null)
  const [error, setError] = useState<string | null>(null)

  const loadData = useCallback(async () => {
    try {
      setIsLoading(true)
      setError(null)
      const [transactionItems, summaryData] = await Promise.all([
        fetchTransactions(filters),
        fetchSummary(),
      ])
      setTransactions(transactionItems ?? [])
      setSummary(summaryData ?? emptySummary)
    } catch (loadError) {
      setError(loadError instanceof Error ? loadError.message : 'Falha ao carregar dados')
    } finally {
      setIsLoading(false)
    }
  }, [filters])

  useEffect(() => {
    void loadData()
  }, [loadData])

  function openCreateModal() {
    setSelectedTransaction(null)
    setIsModalOpen(true)
  }

  function openEditModal(transaction: Transaction) {
    setSelectedTransaction(transaction)
    setIsModalOpen(true)
  }

  function closeModal() {
    setIsModalOpen(false)
    setSelectedTransaction(null)
  }

  async function handleSubmitTransaction(payload: CreateTransactionPayload) {
    try {
      setIsSaving(true)

      if (selectedTransaction) {
        await updateTransaction(selectedTransaction.id, payload)
      } else {
        await createTransaction(payload)
      }

      closeModal()
      await loadData()
    } catch (saveError) {
      setError(saveError instanceof Error ? saveError.message : 'Falha ao salvar transação')
    } finally {
      setIsSaving(false)
    }
  }

  async function handleDeleteTransaction(id: string) {
    try {
      await deleteTransaction(id)
      await loadData()
    } catch (deleteError) {
      setError(deleteError instanceof Error ? deleteError.message : 'Falha ao excluir transação')
    }
  }

  return (
    <main className="page-shell">
      <header className="hero">
        <div>
          <h1>Controle Financeiro</h1>
          <p className="subtitle">
            O controle financeiro é a base para uma vida mais tranquila e para a saúde de qualquer negócio.
          </p>
        </div>
        <button className="primary-button" onClick={openCreateModal}>
          Nova transação
        </button>
      </header>

      {error ? <div className="alert">{error}</div> : null}

      <SummaryCards summary={summary} />
      <FilterBar filters={filters} onChange={setFilters} />

      {isLoading ? (
        <div className="empty-state">Carregando transações...</div>
      ) : (
        <TransactionsTable transactions={transactions} onDelete={handleDeleteTransaction} onEdit={openEditModal} />
      )}

      <TransactionForm
        open={isModalOpen}
        loading={isSaving}
        transaction={selectedTransaction}
        onClose={closeModal}
        onSubmit={handleSubmitTransaction}
      />
    </main>
  )
}
