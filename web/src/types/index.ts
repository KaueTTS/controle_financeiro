export type TransactionType = 'income' | 'expense'

export interface Transaction {
  id: string
  title: string
  description: string
  amount: number
  type: TransactionType
  category: string
  created_at: string
}

export interface Summary {
  income: number
  expense: number
  balance: number
}

export interface ApiResponse<T> {
  data: T
}

export interface ApiError {
  message: string
  errors?: Array<{
    field: string
    message: string
  }>
}

export interface Filters {
  search: string
  type: string
  category: string
}
