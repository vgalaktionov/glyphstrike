const LineLength = 4 + 6 + 6; // max 4 byte rune + 6 byte foreground hex + 6 byte background hex

class CanvasRenderer {
  #backgroundColor;
  #canvas;
  #decoder;
  constructor(backgroundColor) {
    this.#canvas = document.getElementById("screen");
    if (!this.#canvas instanceof HTMLCanvasElement)
      throw new Error("invalid canvas");
    this.#canvas.width = window.innerWidth;
    this.#canvas.height = window.innerHeight;
    this.#canvas.imageSmoothingEnabled = false;

    this.#backgroundColor = backgroundColor;
    document.body.style.backgroundColor = backgroundColor;
    this.#decoder = new TextDecoder();

    const ctx = this.#canvas.getContext("2d");
    ctx.fillStyle = this.#backgroundColor;
    ctx.fillRect(0, 0, this.#canvas.width, this.#canvas.height);

    ctx.font =
      '12pt ui-monospace, Menlo, Monaco, "Cascadia Mono", "Segoe UI Mono", "Roboto Mono", "Oxygen Mono", "Ubuntu Monospace", "Source Code Pro","Fira Mono", "Droid Sans Mono", "Courier New", monospace';
  }

  clear() {
    const ctx = this.#canvas.getContext("2d");
    ctx.fillStyle = this.#backgroundColor;
    ctx.fillRect(0, 0, canvas.width, canvas.height);
  }

  async pollEvent() {
    return new Promise((resolve, reject) => {
      const listenEvent = (ev) => {
        ev.preventDefault();
        if (ev instanceof KeyboardEvent) {
          resolve({ type: "key", key: ev.key });
        } else if (ev instanceof MouseEvent) {
          resolve({
            type: "mouse",
            x: ev.screenX,
            y: ev.screenY,
            button: ev.button,
          });
        }
        window.removeEventListener("keydown", listenEvent);
        window.removeEventListener("mousedown", listenEvent);
        window.removeEventListener("mousemove", listenEvent);
      };
      window.addEventListener("keydown", listenEvent);
      window.removeEventListener("mousedown", listenEvent);
      window.addEventListener("mousemove", listenEvent);
    });
  }

  postKeyEvent(key) {
    window.dispatchEvent(new KeyboardEvent("keydown", { key }));
  }

  postMouseEvent(x, y, button) {
    window.dispatchEvent(
      new MouseEvent("mousedown", { screenX: x, screenY: y, button })
    );
  }

  size() {
    const measure = this.#canvas.getContext("2d").measureText("█");
    return {
      width: Math.floor(window.innerWidth / measure.width),
      height: Math.floor(
        window.innerHeight /
          (measure.actualBoundingBoxAscent +
            measure.actualBoundingBoxDescent -
            2)
      ),
    };
  }

  cellSize() {
    const measure = this.#canvas.getContext("2d").measureText("█");
    return {
      width: measure.width,
      height:
        measure.actualBoundingBoxAscent + measure.actualBoundingBoxDescent - 2,
    };
  }

  show(buf) {
    console.debug("called show");
    let xOffset = 0;
    let yOffset = 0;
    const { width: cellWidth, height: cellHeight } = this.cellSize();
    const ctx = this.#canvas.getContext("2d", { colorSpace: "display-p3" });
    for (let i = 0; i < buf.length; i += LineLength) {
      const char = this.#decoder.decode(buf.slice(i, i + 4)).trimStart();
      const fg = "#" + this.#decoder.decode(buf.slice(i + 4, i + 4 + 6));
      const bg =
        "#" + this.#decoder.decode(buf.slice(i + 4 + 6, i + LineLength));
      ctx.fillStyle = bg;
      // draw slightly more to hide ugly gaps
      ctx.fillRect(
        xOffset,
        yOffset - cellHeight + 3,
        cellWidth,
        cellHeight + 3
      );
      ctx.fillStyle = fg;
      ctx.fillText(char, xOffset, yOffset);

      yOffset += cellHeight;
      if (yOffset + cellHeight >= this.#canvas.height) {
        yOffset = 0;
        xOffset += cellWidth;
      }
    }
  }

  sync() {
    this.#canvas.width = window.innerWidth;
    this.#canvas.height = window.innerHeight;
    this.show();
  }
}

window.CanvasRenderer = CanvasRenderer;
