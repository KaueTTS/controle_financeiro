export function getErrorMessage(error: unknown, fallback = 'Erro inesperado.') {
  return error instanceof Error ? error.message : fallback;
}
