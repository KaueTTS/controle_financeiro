import { useCallback, useEffect, useState } from 'react'
import { FilterBar } from '../components/FilterBar'
import { SummaryCards } from '../components/SummaryCards'
import { TransactionForm } from '../components/TransactionForm'
import { TransactionsTable } from '../components/TransactionsTable'
import {
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
      setError(loadError instanceof Error ? loadError.message : 'Failed to load data')
    } finally {
      setIsLoading(false)
    }
  }, [filters])

  useEffect(() => {
    void loadData()
  }, [loadData])

  async function handleCreateTransaction(payload: CreateTransactionPayload) {
    try {
      setIsSaving(true)
      await createTransaction(payload)
      setIsModalOpen(false)
      await loadData()
    } catch (saveError) {
      setError(saveError instanceof Error ? saveError.message : 'Failed to save transaction')
    } finally {
      setIsSaving(false)
    }
  }

  async function handleDeleteTransaction(id: string) {
    try {
      await deleteTransaction(id)
      await loadData()
    } catch (deleteError) {
      setError(deleteError instanceof Error ? deleteError.message : 'Failed to delete transaction')
    }
  }

  return (
    <main className="page-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">Controle Financeiro</p>
          <h1>Controle Financeiro</h1>
          <p className="subtitle">
            Solução orientada à produção com camadas limpas, validação, filtros e armazenamento persistente.
          </p>
        </div>
        <button className="primary-button" onClick={() => setIsModalOpen(true)}>
          Nova transação
        </button>
      </header>

      {error ? <div className="alert">{error}</div> : null}

      <SummaryCards summary={summary} />
      <FilterBar filters={filters} onChange={setFilters} />

      {isLoading ? (
        <div className="empty-state">Carregando transações...</div>
      ) : (
        <TransactionsTable transactions={transactions} onDelete={handleDeleteTransaction} />
      )}

      <TransactionForm
        open={isModalOpen}
        loading={isSaving}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleCreateTransaction}
      />
    </main>
  )
}
