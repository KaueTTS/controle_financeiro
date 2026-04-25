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
        placeholder="Pesquise por título ou descrição"
      />

      <select
        value={filters.type}
        onChange={(event) => onChange({ ...filters, type: event.target.value })}
      >
        <option value="">Todos os tipos</option>
        <option value="income">Renda</option>
        <option value="expense">Despesa</option>
      </select>

      <input
        value={filters.category}
        onChange={(event) => onChange({ ...filters, category: event.target.value })}
        placeholder="Filtrar por categoria"
      />
    </div>
  )
}
