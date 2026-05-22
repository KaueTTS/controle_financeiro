import { httpClient } from './http';
import type { ApiMessageResponse, PaginatedApiResponse } from '../types/api';
import type { Transaction, TransactionFilters, TransactionPayload } from '../types/transaction';

export async function listTransactions(filters: TransactionFilters) {
  return httpClient<PaginatedApiResponse<Transaction[]>>('/v1/transactions', {
    query: {
      search: filters.search,
      type: filters.type,
      category: filters.category,
      page: filters.page,
      perPage: filters.perPage,
    },
  });
}

export async function createTransaction(payload: TransactionPayload) {
  return httpClient<ApiMessageResponse>('/v1/transactions', {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export async function updateTransaction(id: number, payload: TransactionPayload) {
  return httpClient<ApiMessageResponse>(`/v1/transactions/${id}`, {
    method: 'PUT',
    body: JSON.stringify(payload),
  });
}

export async function deleteTransaction(id: number) {
  return httpClient<void>(`/v1/transactions/${id}`, {
    method: 'DELETE',
  });
}
