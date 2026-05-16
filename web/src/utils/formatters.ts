import type { TransactionType } from '../types/transaction';

export function formatCurrency(value: number) {
  return new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  }).format(value);
}

export function formatDate(value: string) {
  return new Intl.DateTimeFormat('pt-BR', {
    dateStyle: 'short',
    timeStyle: 'short',
  }).format(new Date(value));
}

export function translateTransactionType(type: TransactionType) {
  return type === 'income' ? 'Receita' : 'Despesa';
}
