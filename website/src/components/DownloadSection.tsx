import { useEffect, useRef, useState, type ReactNode } from "react";
import { PLATFORMS, type Platform } from "../config";
import { detectPlatform } from "../detectPlatform";
import PlatformCard from "./PlatformCard";
import { AppleIcon, WindowsIcon, DownloadIcon } from "./Icons";

const ICONS: Record<Platform, ReactNode> = {
  macos: <AppleIcon />,
  windows: <WindowsIcon />,
};

export default function DownloadSection() {
  const [selected, setSelected] = useState<Platform | null>(null);
  const [detected, setDetected] = useState<Platform | null>(null);
  const groupRef = useRef<HTMLDivElement>(null);

  useEffect(() => setDetected(detectPlatform()), []);

  const current = selected ? PLATFORMS[selected] : null;
  const canDownload = !!current && current.url.trim() !== "";

  function handleDownload() {
    if (!current || !canDownload) return;
    // Só inicia o download após o clique de confirmação.
    window.location.assign(current.url);
  }

  function onKeyDown(e: React.KeyboardEvent) {
    if (!["ArrowRight", "ArrowLeft", "ArrowDown", "ArrowUp"].includes(e.key)) return;
    const btns = Array.from(groupRef.current?.querySelectorAll<HTMLButtonElement>("button") ?? []);
    const i = btns.findIndex((b) => b === document.activeElement);
    if (i < 0) return;
    e.preventDefault();
    const next = (i + (e.key === "ArrowRight" || e.key === "ArrowDown" ? 1 : -1) + btns.length) % btns.length;
    btns[next].focus();
  }

  return (
    <section className="download" id="download" aria-labelledby="download-title">
      <div className="container">
        <h2 id="download-title" className="section__title">Baixe o Refletunes</h2>
        <p className="section__lead">
          Selecione seu sistema operacional para baixar a versão compatível com seu computador.
          O Refletunes espelha celulares Android e é instalado no computador — se você está
          acessando pelo celular, abra este site no PC ou Mac.
        </p>

        {detected && (
          <p className="detect" role="status">
            <span className="detect__dot" aria-hidden />
            Detectamos que você está usando {PLATFORMS[detected].name}.
          </p>
        )}

        <div
          className="download__cards"
          role="radiogroup"
          aria-label="Sistema operacional"
          ref={groupRef}
          onKeyDown={onKeyDown}
        >
          {(Object.keys(PLATFORMS) as Platform[]).map((id) => (
            <PlatformCard
              key={id}
              id={id}
              name={PLATFORMS[id].name}
              description={PLATFORMS[id].description}
              icon={ICONS[id]}
              selected={selected === id}
              detected={detected === id}
              available={PLATFORMS[id].url.trim() !== ""}
              onSelect={setSelected}
            />
          ))}
        </div>

        <div className="download__action" aria-live="polite">
          {current && (
            <>
              <button
                type="button"
                className="btn btn--primary btn--lg"
                disabled={!canDownload}
                onClick={handleDownload}
              >
                <DownloadIcon />
                Baixar para {current.name}
              </button>
              {canDownload && (
                <p className="download__note">
                  Versão beta 0.1.0 · arquivo .zip portátil — extraia e execute o Refletunes.
                </p>
              )}
              {!canDownload && <p className="download__note">Esta versão ainda não está disponível.</p>}
            </>
          )}
          {!current && <p className="download__hint">Escolha um sistema para continuar.</p>}
        </div>
      </div>
    </section>
  );
}
