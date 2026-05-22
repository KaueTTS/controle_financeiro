export interface ApiDataResponse<T> {
  data: T;
}

export interface Pagination {
  page: number;
  perPage: number;
  pageCount: number;
  total: number;
}

export interface PaginatedApiResponse<T> extends ApiDataResponse<T> {
  pagination: Pagination;
}

export interface ApiMessageResponse {
  message: string;
}

export interface ApiErrorResponse {
  error?: string;
  message?: string;
}
