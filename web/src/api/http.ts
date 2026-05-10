import type { ApiErrorResponse } from '../types/api';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';
const DEFAULT_API_ERROR_MESSAGE = 'Erro inesperado ao comunicar com a API.';

interface RequestOptions extends RequestInit {
  query?: Record<string, string | number | boolean | undefined | null>;
}

function buildUrl(path: string, query?: RequestOptions['query']) {
  const url = new URL(path, API_BASE_URL);

  Object.entries(query ?? {}).forEach(([key, value]) => {
    if (value !== undefined && value !== null && value !== '') {
      url.searchParams.set(key, String(value));
    }
  });

  return url.toString();
}

function parseJsonSafely(text: string) {
  try {
    return { data: JSON.parse(text) as unknown, success: true };
  } catch {
    return { data: null, success: false };
  }
}

function getErrorMessage(data: unknown, fallback: string) {
  if (typeof data === 'string' && data.trim()) {
    return data.trim();
  }

  if (data && typeof data === 'object') {
    const apiError = data as ApiErrorResponse;
    return apiError.error ?? apiError.message ?? fallback;
  }

  return fallback;
}

function createHeaders(headers?: HeadersInit, body?: BodyInit | null) {
  const requestHeaders = new Headers(headers);
  const shouldSetJsonContentType = body !== undefined && body !== null && !(body instanceof FormData);

  if (shouldSetJsonContentType && !requestHeaders.has('Content-Type')) {
    requestHeaders.set('Content-Type', 'application/json');
  }

  return requestHeaders;
}

async function parseResponse<T>(response: Response): Promise<T> {
  if (response.status === 204 || response.status === 205) {
    return undefined as T;
  }

  const text = await response.text();
  const hasBody = text.trim().length > 0;
  const parsedJson = hasBody ? parseJsonSafely(text) : { data: null, success: false };
  const data = hasBody && parsedJson.success ? parsedJson.data : hasBody ? text : null;

  if (!response.ok) {
    const fallback = `${DEFAULT_API_ERROR_MESSAGE} Status: ${response.status}.`;
    throw new Error(getErrorMessage(data, fallback));
  }

  if (!hasBody) {
    return undefined as T;
  }

  return data as T;
}

export async function httpClient<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const { query, headers, body, ...requestOptions } = options;

  const response = await fetch(buildUrl(path, query), {
    ...requestOptions,
    headers: createHeaders(headers, body),
    body,
  });

  return parseResponse<T>(response);
}
