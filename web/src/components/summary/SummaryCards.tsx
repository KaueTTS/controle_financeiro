import type { Summary } from '../../types/summary';
import { formatCurrency } from '../../utils/formatters';

interface SummaryCardsProps {
  summary: Summary;
}

type SummaryCardVariant = 'income' | 'expense';

interface SummaryCardItem {
  label: string;
  value: number;
  variant: SummaryCardVariant;
  description: string;
  percentage: number;
}

function clampPercentage(value: number) {
  return Math.min(Math.max(value, 0), 100);
}

function TrendIcon({ variant }: { variant: SummaryCardVariant }) {
  return (
    <svg aria-hidden="true" viewBox="0 0 24 24" className="summary-icon-svg">
      <path d="M4 19V5" />
      <path d="M4 19h16" />
      {variant === 'income' ? (
        <>
          <path d="m7 15 4-4 3 3 5-7" />
          <path d="M16 7h3v3" />
        </>
      ) : (
        <>
          <path d="m7 9 4 4 3-3 5 6" />
          <path d="M16 16h3v-3" />
        </>
      )}
    </svg>
  );
}

function SummaryCard({ item }: { item: SummaryCardItem }) {
  const progressWidth = `${clampPercentage(item.percentage)}%`;

  return (
    <article className={`summary-card ${item.variant}`}>
      <div className="summary-card-header">
        <span className="summary-card-icon" aria-hidden="true">
          <TrendIcon variant={item.variant} />
        </span>

        <div>
          <span className="summary-card-kicker">{item.label}</span>
          <p>{item.description}</p>
        </div>
      </div>

      <strong>{formatCurrency(item.value)}</strong>

      <div className="summary-progress" aria-label={`${item.label}: ${item.percentage.toFixed(0)}% do movimento total`}>
        <span style={{ width: progressWidth }} />
      </div>

      <small>{item.percentage.toFixed(0)}% do movimento total</small>
    </article>
  );
}

export function SummaryCards({ summary }: SummaryCardsProps) {
  const totalMovement = summary.income + summary.expense;
  const incomePercentage = totalMovement > 0 ? (summary.income / totalMovement) * 100 : 0;
  const expensePercentage = totalMovement > 0 ? (summary.expense / totalMovement) * 100 : 0;

  const cards: SummaryCardItem[] = [
    {
      label: 'Receitas',
      value: summary.income,
      variant: 'income',
      description: 'Total de entradas no período',
      percentage: incomePercentage,
    },
    {
      label: 'Despesas',
      value: summary.expense,
      variant: 'expense',
      description: 'Total de saídas no período',
      percentage: expensePercentage,
    },
  ];

  return (
    <section className="summary-panel" aria-label="Resumo financeiro">
      <div className="summary-panel-header">
        <span className="eyebrow">Movimentação</span>
        <p>Comparativo entre entradas e saídas.</p>
      </div>

      <div className="summary-chart" aria-hidden="true">
        <span className="summary-chart-income" style={{ flex: Math.max(summary.income, 1) }} />
        <span className="summary-chart-expense" style={{ flex: Math.max(summary.expense, 1) }} />
      </div>

      <div className="summary-grid">
        {cards.map((item) => (
          <SummaryCard item={item} key={item.label} />
        ))}
      </div>
    </section>
  );
}
