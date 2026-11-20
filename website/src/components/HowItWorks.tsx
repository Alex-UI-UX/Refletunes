import GlassIcons from "./GlassIcons.jsx";
import { DownloadIcon, PhoneIcon, ScreenIcon } from "./Icons";

const STEPS = [
  { n: "01", icon: <DownloadIcon />, color: "blue", title: "Baixe o Refletunes", text: "Extraia o arquivo baixado e abra o Refletunes no seu computador." },
  { n: "02", icon: <PhoneIcon />, color: "purple", title: "Conecte seu Android", text: "Cabo USB ou Wi-Fi. O app orienta a ativação da depuração no celular." },
  { n: "03", icon: <ScreenIcon />, color: "teal", title: "Espelhe sua tela", text: "Visualize e controle a tela do celular com mouse e teclado do computador." },
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
                <GlassIcons
                  items={[{ icon: s.icon, color: s.color, label: s.title }]}
                  size={44}
                  gap={0}
                  columns={1}
                  plate="tilt"
                  tilt={12}
                  labels="none"
                  intro={false}
                  theme="dark"
                  className="step__icon"
                />
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
