try {
  window.eruda.init();
} catch (error) {}

window.initializeScreen = function (blackColor) {
  const canvas = document.createElement("canvas");
  canvas.width = window.outerWidth;
  canvas.height = window.outerHeight;
  document.body.appendChild(canvas);

  const ctx = canvas.getContext("2d");
  ctx.fillStyle = blackColor;
  ctx.fillRect(0, 0, canvas.width, canvas.height);

  ctx.font =
    '12pt ui-monospace, Menlo, Monaco, "Cascadia Mono", "Segoe UI Mono", "Roboto Mono", "Oxygen Mono", "Ubuntu Monospace", "Source Code Pro","Fira Mono", "Droid Sans Mono", "Courier New", monospace';

  window.ctx = ctx;
  window.blackColor = blackColor;
};

window.clear = function () {
  if (window.ctx != null && window.ctx instanceof CanvasRenderingContext2D) {
    window.ctx.fillStyle = window.blackColor;
    window.ctx.fillRect(0, 0, canvas.width, canvas.height);
  }
};

window.pollEvent = async function () {
  return new Promise((resolve, reject) => {
    const listenEvent = (ev) => {
      console.debug("calling pollEvent: ", ev);
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
};

window.postKeyEvent = async function (key) {
  console.debug("calling postKeyEvent");
  window.dispatchEvent(new KeyboardEvent("keydown", { key }));
};

window.postMouseEvent = function (x, y, button) {
  console.debug("calling postMouseEvent");
  window.dispatchEvent(
    new MouseEvent("mousedown", { screenX: x, screenY: y, button })
  );
};

window.size = function () {
  let result = { width: 0, height: 0 };
  if (window.ctx != null && window.ctx instanceof CanvasRenderingContext2D) {
    const measure = ctx.measureText("█");
    result = {
      width: Math.floor(window.innerWidth / measure.width),
      height: Math.floor(
        window.innerHeight /
          (measure.actualBoundingBoxAscent +
            measure.actualBoundingBoxDescent -
            2)
      ),
    };
  }
  return result;
};

window.cellSize = function () {
  let result = { width: 1, height: 1 };
  if (window.ctx != null && window.ctx instanceof CanvasRenderingContext2D) {
    const measure = ctx.measureText("█");
    result = {
      width: measure.width,
      height:
        measure.actualBoundingBoxAscent + measure.actualBoundingBoxDescent - 2,
    };
  }
  return result;
};

window.show = function (buf) {
  if (window.ctx != null && window.ctx instanceof CanvasRenderingContext2D) {
    let xOffset = 0;
    let yOffset = 0;
    const { width: cellWidth, height: cellHeight } = window.cellSize();
    for (const column of JSON.parse(buf)) {
      for (const cell of column) {
        window.ctx.fillStyle = cell.background;
        window.ctx.fillRect(xOffset, yOffset, cellWidth, cellHeight);
        window.ctx.fillStyle = cell.foreground;
        window.ctx.fillText(cell.text, xOffset, yOffset);
        yOffset += cellHeight;
      }
      xOffset += cellWidth;
      yOffset = 0;
    }
  }
};
