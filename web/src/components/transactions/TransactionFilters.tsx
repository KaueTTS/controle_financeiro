import type { TransactionFilters as Filters } from '../../types/transaction';

interface TransactionFiltersProps {
  filters: Filters;
  categories: string[];
  hasActiveFilters: boolean;
  onChange: (filters: Filters) => void;
  onClear: () => void;
}

export function TransactionFilters({
  filters,
  categories,
  hasActiveFilters,
  onChange,
  onClear,
}: TransactionFiltersProps) {
  return (
    <section className="filters-panel">
      <div className="section-title">
        <div>
          <span className="eyebrow">Filtros</span>
        </div>

        {hasActiveFilters && (
          <button className="ghost-button small-button" type="button" onClick={onClear}>
            Limpar filtros
          </button>
        )}
      </div>

      <div className="filters-grid">
        <div className="field search-field">
          <label htmlFor="search">Buscar</label>
          <input
            id="search"
            type="text"
            placeholder="Título ou descrição"
            value={filters.search}
            onChange={(event) => onChange({ ...filters, search: event.target.value, page: 1 })}
          />
        </div>

        <div className="field">
          <label htmlFor="type">Tipo</label>
          <select
            id="type"
            value={filters.type}
            onChange={(event) => onChange({ ...filters, type: event.target.value as Filters['type'], page: 1 })}
          >
            <option value="">Todos</option>
            <option value="income">Receitas</option>
            <option value="expense">Despesas</option>
          </select>
        </div>

        <div className="field">
          <label htmlFor="category">Categoria</label>
          <select
            id="category"
            value={filters.category}
            onChange={(event) => onChange({ ...filters, category: event.target.value, page: 1 })}
          >
            <option value="">Todas</option>
            {categories.map((category) => (
              <option key={category} value={category}>
                {category}
              </option>
            ))}
          </select>
        </div>
      </div>
    </section>
  );
}
