import type { ReactNode } from "react";
import { CheckIcon } from "./Icons";
import type { Platform } from "../config";

type Props = {
  id: Platform;
  name: string;
  description: string;
  icon: ReactNode;
  selected: boolean;
  detected: boolean;
  available: boolean;
  onSelect: (id: Platform) => void;
};

export default function PlatformCard({ id, name, description, icon, selected, detected, available, onSelect }: Props) {
  return (
    <button
      type="button"
      role="radio"
      aria-checked={selected}
      className={`pcard${selected ? " is-selected" : ""}`}
      onClick={() => onSelect(id)}
    >
      <span className="pcard__check" aria-hidden>
        <CheckIcon width={14} height={14} />
      </span>
      <span className="pcard__icon">{icon}</span>
      <span className="pcard__name">{name}</span>
      <span className="pcard__desc">{description}</span>
      {detected && <span className="pcard__tag">Sugerido para você</span>}
      {!available && <span className="pcard__unavail">Esta versão ainda não está disponível.</span>}
      <span className="pcard__select">{selected ? `${name} selecionado` : `Selecionar ${name}`}</span>
    </button>
  );
}
