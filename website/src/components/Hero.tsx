import ProductDemo from "./ProductDemo";

export default function Hero() {
  return (
    <section className="hero" id="topo">
      <div className="container hero__inner">
        <div className="hero__copy">
          <h1 className="hero__title">Seu celular. Na tela do seu computador.</h1>
          <p className="hero__sub">
            Espelhe a tela do seu smartphone diretamente no computador, com uma experiência simples, rápida e sem complicação.
          </p>
          <div className="hero__actions">
            <a href="#download" className="btn btn--primary">Baixar Refletunes</a>
            <a href="#como-funciona" className="btn btn--ghost">Como funciona</a>
          </div>
          <p className="hero__meta">Para celulares Android · Gratuito</p>
        </div>
        <ProductDemo />
      </div>
    </section>
  );
}
