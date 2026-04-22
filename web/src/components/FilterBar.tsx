import type { Filters } from '../types'

interface FilterBarProps {
  filters: Filters
  onChange: (filters: Filters) => void
}

export function FilterBar({ filters, onChange }: FilterBarProps) {
  return (
    <div className="toolbar filters-grid">
      <input
        value={filters.search}
        onChange={(event) => onChange({ ...filters, search: event.target.value })}
        placeholder="Search by title or category"
      />

      <select
        value={filters.type}
        onChange={(event) => onChange({ ...filters, type: event.target.value })}
      >
        <option value="">All types</option>
        <option value="income">Income</option>
        <option value="expense">Expense</option>
      </select>

      <input
        value={filters.category}
        onChange={(event) => onChange({ ...filters, category: event.target.value })}
        placeholder="Filter by category"
      />
    </div>
  )
}
