import { GITHUB_URL } from "../config";

export function Logo() {
  return <span className="logo">Refletunes</span>;
}

export default function Header() {
  return (
    <header className="header">
      <div className="container header__inner">
        <a href="#topo" className="header__brand" aria-label="Refletunes, início">
          <Logo />
        </a>
        <nav className="header__nav" aria-label="Principal">
          <a href="#como-funciona" className="header__link">Como funciona</a>
          <a href="#download" className="header__link header__link--cta">Download</a>
          {GITHUB_URL && (
            <a href={GITHUB_URL} className="header__link" target="_blank" rel="noopener noreferrer">GitHub</a>
          )}
        </nav>
      </div>
    </header>
  );
}
