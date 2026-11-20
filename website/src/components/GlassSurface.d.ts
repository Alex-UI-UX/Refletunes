import type { CSSProperties, ReactNode, RefObject } from "react";

declare module "./GlassSurface.jsx" {
  export interface GlassSurfaceProps {
    children?: ReactNode;
    width?: number | string;
    height?: number | string;
    borderRadius?: number;
    borderWidth?: number;
    brightness?: number;
    opacity?: number;
    blur?: number;
    displace?: number;
    backgroundOpacity?: number;
    saturation?: number;
    distortionScale?: number;
    redOffset?: number;
    greenOffset?: number;
    blueOffset?: number;
    xChannel?: "R" | "G" | "B";
    yChannel?: "R" | "G" | "B";
    mixBlendMode?: string;
    refraction?: boolean;
    renderer?: "auto" | "svg" | "webgl";
    backdrop?: RefObject<HTMLImageElement | HTMLVideoElement | HTMLCanvasElement | null>;
    shape?: string;
    className?: string;
    style?: CSSProperties;
  }
  const GlassSurface: (props: GlassSurfaceProps) => JSX.Element;
  export default GlassSurface;
}
