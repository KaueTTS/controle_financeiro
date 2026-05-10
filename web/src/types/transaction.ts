export type TransactionType = 'income' | 'expense';

export interface Transaction {
  id: number;
  title: string;
  description?: string | null;
  amount: number;
  category: string;
  type: TransactionType;
  created_at: string;
}

export interface TransactionPayload {
  title: string;
  description?: string | null;
  amount: number;
  category: string;
  type: TransactionType;
}

export interface TransactionFilters {
  search: string;
  type: '' | TransactionType;
  category: string;
}
