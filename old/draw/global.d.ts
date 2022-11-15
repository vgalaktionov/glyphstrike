export {};

declare global {
  interface CanvasRenderingContext2D {
    imageSmoothingEnabled: boolean;
  }

  interface Window {
    CanvasRenderer: new (backgroundColor: string) => CanvasRenderer;
  }
}
