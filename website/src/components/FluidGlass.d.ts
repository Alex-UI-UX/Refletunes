import type { CSSProperties, ReactNode } from "react";

declare module "./FluidGlass.jsx" {
  export interface FluidGlassProps {
    text?: string;
    textColor?: string;
    fontFamily?: string;
    fontWeight?: number | string;
    fontSize?: number;
    src?: string;
    fit?: "cover" | "contain";
    background?: string;
    shape?: "torus" | "lens" | "blob" | "pill" | "cube" | "prism" | "diamond";
    shapes?: string[];
    morphOnClick?: boolean;
    onShapeChange?: (shape: string) => void;
    size?: number;
    refraction?: number;
    dispersion?: number;
    reflection?: number;
    frost?: number;
    color?: string;
    tint?: number;
    shadow?: number;
    elevation?: number;
    interaction?: "follow" | "drag" | "static";
    smoothness?: number;
    wobble?: number;
    spin?: number;
    idle?: boolean;
    intro?: boolean;
    position?: [number, number];
    maxDevicePixelRatio?: number;
    className?: string;
    style?: CSSProperties;
    children?: ReactNode;
  }
  const FluidGlass: (props: FluidGlassProps) => JSX.Element;
  export default FluidGlass;
}
