import FluidGlass from "./FluidGlass.jsx";

/** Tela genérica de smartphone (apps fictícios sem marca), usada no celular e espelhada no monitor. */
function PhoneScreen() {
  const tiles = [
    "#386BFF", "#2a3350", "#A12BFF", "#2a3350",
    "#1d2540", "#00C8E8", "#2a3350", "#386BFF",
    "#2a3350", "#A12BFF", "#1d2540", "#2a3350",
  ];
  return (
    <div className="pscreen">
      <div className="pscreen__time">9:41</div>
      <div className="pscreen__date">sexta-feira, 20 de novembro</div>
      <div className="pscreen__grid">
        {tiles.map((c, i) => (
          <span key={i} style={{ background: c, opacity: c === "#2a3350" || c === "#1d2540" ? 1 : 0.85 }} />
        ))}
      </div>
      <div className="pscreen__dock">
        <span /><span /><span /><span />
      </div>
    </div>
  );
}

export default function ProductDemo() {
  return (
    <div className="demo" role="img" aria-label="Um smartphone transmitindo sua tela para a janela do Refletunes em um monitor">
      <div className="monitor">
        <div className="monitor__bezel">
          <div className="monitor__screen">
            <div className="appwin">
              <div className="appwin__bar">
                <span className="appwin__dots"><i /><i /><i /></span>
                <span className="appwin__title">Refletunes</span>
              </div>
              <div className="appwin__body">
                <div className="mirror">
                  <PhoneScreen />
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="monitor__neck" />
        <div className="monitor__base" />
      </div>

      <div className="link" aria-hidden>
        <span className="link__line" />
        <span className="link__dot" />
        <span className="link__dot link__dot--2" />
        <span className="link__dot link__dot--3" />
      </div>

      <div className="phone">
        <div className="phone__island" />
        <PhoneScreen />
      </div>

      <FluidGlass
        text=""
        shape="lens"
        size={190}
        interaction="follow"
        refraction={1}
        dispersion={0.25}
        reflection={0.8}
        shadow={0.5}
        tint={0.15}
        color="#bcd6ff"
        smoothness={0.6}
        morphOnClick={false}
        idle={false}
        className="demo__glass"
      />
    </div>
  );
}
