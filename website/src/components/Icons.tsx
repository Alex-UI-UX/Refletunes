import type { SVGProps } from "react";

type P = SVGProps<SVGSVGElement>;
const base = { width: 20, height: 20, viewBox: "0 0 24 24", fill: "none", stroke: "currentColor", strokeWidth: 1.6, strokeLinecap: "round" as const, strokeLinejoin: "round" as const, "aria-hidden": true };

export const AppleIcon = (p: P) => (
  <svg viewBox="0 0 24 24" width={28} height={28} fill="currentColor" aria-hidden {...p}>
    <path d="M16.37 12.64c-.03-2.6 2.12-3.85 2.22-3.91-1.21-1.77-3.09-2.01-3.76-2.04-1.6-.16-3.12.94-3.93.94-.81 0-2.06-.92-3.39-.89-1.74.03-3.35 1.01-4.25 2.57-1.81 3.14-.46 7.78 1.3 10.33.86 1.25 1.89 2.65 3.24 2.6 1.3-.05 1.79-.84 3.37-.84 1.57 0 2.02.84 3.4.81 1.4-.02 2.29-1.27 3.14-2.52.99-1.45 1.4-2.85 1.42-2.92-.03-.01-2.73-1.05-2.76-4.13zM13.79 4.96c.72-.87 1.2-2.08 1.07-3.28-1.03.04-2.28.69-3.02 1.55-.66.77-1.24 2-1.09 3.18 1.15.09 2.32-.58 3.04-1.45z" />
  </svg>
);

export const WindowsIcon = (p: P) => (
  <svg viewBox="0 0 24 24" width={26} height={26} fill="currentColor" aria-hidden {...p}>
    <path d="M3 5.55 10.5 4.5v6.9H3zM11.5 4.36 21 3v8.4h-9.5zM3 12.6h7.5v6.9L3 18.45zM11.5 12.6H21V21l-9.5-1.34z" />
  </svg>
);

export const DownloadIcon = (p: P) => (
  <svg {...base} {...p}><path d="M12 4v11m0 0-4-4m4 4 4-4M5 19h14" /></svg>
);
export const PhoneIcon = (p: P) => (
  <svg {...base} {...p}><rect x="7" y="3" width="10" height="18" rx="2.2" /><path d="M11 18h2" /></svg>
);
export const ScreenIcon = (p: P) => (
  <svg {...base} {...p}><rect x="3" y="4" width="18" height="12" rx="2" /><path d="M9 20h6M12 16v4" /></svg>
);
export const CheckIcon = (p: P) => (
  <svg {...base} strokeWidth={2.2} {...p}><path d="m5 12.5 4.5 4.5L19 7.5" /></svg>
);
export const GithubIcon = (p: P) => (
  <svg viewBox="0 0 24 24" width={18} height={18} fill="currentColor" aria-hidden {...p}>
    <path d="M12 2a10 10 0 0 0-3.16 19.49c.5.09.68-.22.68-.48v-1.7c-2.78.6-3.37-1.34-3.37-1.34-.45-1.15-1.11-1.46-1.11-1.46-.91-.62.07-.61.07-.61 1 .07 1.53 1.03 1.53 1.03.9 1.52 2.34 1.08 2.91.83.09-.65.35-1.08.63-1.33-2.22-.25-4.55-1.11-4.55-4.94 0-1.09.39-1.98 1.03-2.68-.1-.25-.45-1.27.1-2.64 0 0 .84-.27 2.75 1.02a9.5 9.5 0 0 1 5 0c1.91-1.29 2.75-1.02 2.75-1.02.55 1.37.2 2.39.1 2.64.64.7 1.03 1.59 1.03 2.68 0 3.84-2.34 4.69-4.57 4.93.36.31.68.92.68 1.85v2.74c0 .27.18.58.69.48A10 10 0 0 0 12 2z" />
  </svg>
);
