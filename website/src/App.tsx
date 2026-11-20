import Header from "./components/Header";
import Hero from "./components/Hero";
import DownloadSection from "./components/DownloadSection";
import HowItWorks from "./components/HowItWorks";
import Footer from "./components/Footer";

export default function App() {
  return (
    <>
      <a className="skip" href="#conteudo">Ir para o conteúdo</a>
      <Header />
      <main id="conteudo">
        <Hero />
        <DownloadSection />
        <HowItWorks />
      </main>
      <Footer />
    </>
  );
}
