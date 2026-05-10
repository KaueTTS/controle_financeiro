import { httpClient } from './http';
import type { ApiDataResponse } from '../types/api';
import type { Summary } from '../types/summary';
import type { TransactionFilters } from '../types/transaction';

export async function getSummary(filters: TransactionFilters) {
  const response = await httpClient<ApiDataResponse<Summary>>('/v1/summary', {
    query: {
      search: filters.search,
      type: filters.type,
      category: filters.category,
    },
  });

  return response.data;
}
