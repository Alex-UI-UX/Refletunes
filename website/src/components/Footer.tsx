import { GITHUB_URL } from "../config";
import { Logo } from "./Header";

export default function Footer() {
  return (
    <footer className="footer">
      <div className="container footer__inner">
        <div className="footer__brand">
          <Logo />
          <p>Uma maneira simples de conectar seu smartphone ao computador.</p>
        </div>
        <nav className="footer__nav" aria-label="Rodapé">
          <a href="#download">Download</a>
          <a href="#como-funciona">Como funciona</a>
          {GITHUB_URL && <a href={GITHUB_URL} target="_blank" rel="noopener noreferrer">GitHub</a>}
        </nav>
      </div>
      <div className="container footer__copy">© {new Date().getFullYear()} Refletunes</div>
    </footer>
  );
}
