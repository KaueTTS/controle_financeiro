import type { Summary } from '../types'

interface SummaryCardsProps {
  summary: Summary
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(value)
}

export function SummaryCards({ summary }: SummaryCardsProps) {
  const cards = [
    { label: 'Income', value: summary.income },
    { label: 'Expenses', value: summary.expense },
    { label: 'Balance', value: summary.balance },
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
