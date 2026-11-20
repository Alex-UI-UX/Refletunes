import type { Platform } from "./config";

/** Sugere o sistema do visitante. Nunca seleciona nem baixa nada. */
export function detectPlatform(): Platform | null {
  if (typeof navigator === "undefined") return null;
  const nav = navigator as Navigator & { userAgentData?: { platform?: string } };
  const hint = `${nav.userAgentData?.platform ?? ""} ${nav.platform ?? ""} ${nav.userAgent ?? ""}`;
  // iPhone/iPad não são computadores macOS.
  if (/iPhone|iPad|iPod|Android/i.test(hint)) return null;
  if (/Mac/i.test(hint)) return "macos";
  if (/Win/i.test(hint)) return "windows";
  return null;
}
