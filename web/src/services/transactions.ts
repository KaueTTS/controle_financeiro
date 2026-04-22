import { api } from './api'
import type { ApiResponse, Filters, Summary, Transaction, TransactionType } from '../types'

export interface CreateTransactionPayload {
  title: string
  description: string
  amount: number
  type: TransactionType
  category: string
}

export async function fetchTransactions(filters: Filters) {
  const params = new URLSearchParams()

  if (filters.search) params.set('search', filters.search)
  if (filters.type) params.set('type', filters.type)
  if (filters.category) params.set('category', filters.category)

  const response = await api.get<ApiResponse<Transaction[]>>(`/transactions?${params.toString()}`)
  return response.data
}

export async function fetchSummary() {
  const response = await api.get<ApiResponse<Summary>>('/summary')
  return response.data
}

export async function createTransaction(payload: CreateTransactionPayload) {
  await api.post<{ message: string }>('/transactions', payload)
}

export async function deleteTransaction(id: string) {
  await api.delete<void>(`/transactions/${id}`)
}
