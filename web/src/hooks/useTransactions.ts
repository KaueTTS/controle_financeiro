import { useCallback, useEffect, useMemo, useState } from 'react';
import { getSummary } from '../api/summary';
import { createTransaction, deleteTransaction, listTransactions, updateTransaction } from '../api/transactions';
import type { Summary } from '../types/summary';
import type { Transaction, TransactionFilters, TransactionPayload } from '../types/transaction';
import { getErrorMessage } from '../utils/errors';

const INITIAL_FILTERS: TransactionFilters = {
  search: '',
  type: '',
  category: '',
};

const EMPTY_SUMMARY: Summary = {
  income: 0,
  expense: 0,
  balance: 0,
};

export function useTransactions() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [summary, setSummary] = useState<Summary>(EMPTY_SUMMARY);
  const [filters, setFilters] = useState<TransactionFilters>(INITIAL_FILTERS);
  const [isLoading, setIsLoading] = useState(true);
  const [isDeleting, setIsDeleting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const categories = useMemo(() => {
    return Array.from(new Set(transactions.map((transaction) => transaction.category))).sort();
  }, [transactions]);

  const hasActiveFilters = Boolean(filters.search || filters.type || filters.category);

  const refresh = useCallback(async () => {
    try {
      setIsLoading(true);
      setError(null);

      const [transactionsResponse, summaryResponse] = await Promise.all([
        listTransactions(filters),
        getSummary(filters),
      ]);

      setTransactions(transactionsResponse);
      setSummary(summaryResponse);
    } catch (err) {
      setError(getErrorMessage(err));
    } finally {
      setIsLoading(false);
    }
  }, [filters]);

  async function saveTransaction(payload: TransactionPayload, id?: number) {
    try {
      setError(null);

      if (id) {
        await updateTransaction(id, payload);
      } else {
        await createTransaction(payload);
      }

      await refresh();
    } catch (err) {
      setError(getErrorMessage(err, 'Erro ao salvar lançamento.'));
      throw err;
    }
  }

  async function removeTransaction(id: number) {
    try {
      setIsDeleting(true);
      setError(null);
      await deleteTransaction(id);
      await refresh();
    } catch (err) {
      setError(getErrorMessage(err, 'Erro ao excluir lançamento.'));
      throw err;
    } finally {
      setIsDeleting(false);
    }
  }

  function clearFilters() {
    setFilters(INITIAL_FILTERS);
  }

  useEffect(() => {
    void refresh();
  }, [refresh]);

  return {
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
    refresh,
  };
}
