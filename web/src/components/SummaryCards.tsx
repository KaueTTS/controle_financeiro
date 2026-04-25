import type { Summary } from '../types'

interface SummaryCardsProps {
  summary: Summary
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  }).format(value)
}

export function SummaryCards({ summary }: SummaryCardsProps) {
  const cards = [
    { label: 'Renda', value: summary.income },
    { label: 'Despesa', value: summary.expense },
    { label: 'Saldo', value: summary.balance },
  ]

  return (
    <section className="summary-grid">
      {cards.map((card) => (
        <article key={card.label} className="card">
          <span>{card.label}</span>
          <strong>{formatCurrency(card.value)}</strong>
        </article>
      ))}
    </section>
  )
}
