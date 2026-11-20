/**
 * Configuração central do site.
 * Preencha as URLs abaixo quando os instaladores estiverem publicados.
 * Enquanto estiverem vazias, a plataforma aparece como indisponível.
 */
export const DOWNLOAD_MACOS_URL = "";
export const DOWNLOAD_WINDOWS_URL =
  "https://github.com/Alex-UI-UX/Refletunes/releases/download/v0.1.0-beta/refletunes-win64.zip";

/** Deixe vazio para ocultar os links do GitHub. */
export const GITHUB_URL = "https://github.com/Alex-UI-UX/Refletunes";

export type Platform = "macos" | "windows" | "ios";

export const PLATFORMS: Record<
  Platform,
  { id: Platform; name: string; description: string; url: string; comingSoon?: boolean }
> = {
  macos: {
    id: "macos",
    name: "macOS",
    description: "Para iMac e MacBook.",
    url: DOWNLOAD_MACOS_URL,
  },
  windows: {
    id: "windows",
    name: "Windows",
    description: "Para computadores com Windows 10 e 11.",
    url: DOWNLOAD_WINDOWS_URL,
  },
  ios: {
    id: "ios",
    name: "iOS",
    description: "Para iPhone.",
    url: "",
    comingSoon: true,
  },
};
