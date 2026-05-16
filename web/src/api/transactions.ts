import { httpClient } from './http';
import type { ApiDataResponse, ApiMessageResponse } from '../types/api';
import type { Transaction, TransactionFilters, TransactionPayload } from '../types/transaction';

export async function listTransactions(filters: TransactionFilters) {
  const response = await httpClient<ApiDataResponse<Transaction[]>>('/v1/transactions', {
    query: {
      search: filters.search,
      type: filters.type,
      category: filters.category,
    },
  });

  return response.data;
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
