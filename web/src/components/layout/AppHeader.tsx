import type { Theme } from '../../types/theme';
import { ThemeToggle } from './ThemeToggle';

interface AppHeaderProps {
  theme: Theme;
  onToggleTheme: () => void;
  onCreateTransaction: () => void;
}

export function AppHeader({ theme, onToggleTheme, onCreateTransaction }: AppHeaderProps) {
  return (
    <header className="topbar">
      <div className="topbar-copy">
        <span className="eyebrow">Controle financeiro</span>
        <h1>Dashboard financeiro</h1>
        <p>Acompanhe seu saldo, receitas, despesas e histórico.</p>
      </div>

      <div className="topbar-actions">
        <ThemeToggle theme={theme} onToggle={onToggleTheme} />

        <button className="primary-button" type="button" onClick={onCreateTransaction}>
          Nova transação
        </button>
      </div>
    </header>
  );
}
