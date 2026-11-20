import { DownloadIcon, PhoneIcon, ScreenIcon } from "./Icons";

const STEPS = [
  { n: "01", icon: <DownloadIcon />, title: "Baixe o Refletunes", text: "Instale o aplicativo no seu computador." },
  { n: "02", icon: <PhoneIcon />, title: "Conecte seu smartphone", text: "Prepare seu dispositivo para a conexão com o computador." },
  { n: "03", icon: <ScreenIcon />, title: "Espelhe sua tela", text: "Visualize e controle a tela do smartphone diretamente no computador, conforme os recursos disponíveis no aplicativo." },
];

export default function HowItWorks() {
  return (
    <section className="how" id="como-funciona" aria-labelledby="how-title">
      <div className="container">
        <h2 id="how-title" className="section__title">Como funciona</h2>
        <ol className="steps">
          {STEPS.map((s) => (
            <li className="step" key={s.n}>
              <div className="step__head">
                <span className="step__icon">{s.icon}</span>
                <span className="step__n">{s.n}</span>
              </div>
              <h3 className="step__title">{s.title}</h3>
              <p className="step__text">{s.text}</p>
            </li>
          ))}
        </ol>
      </div>
    </section>
  );
}
