//go:build js
// +build js

package draw

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// CanvasBuffer is the internal buffer of js objects to be drawn with the next CanvasRenderer.Show() call.
type CanvasBuffer [][]map[string]interface{}

func NewCanvasBuffer(width, height int) CanvasBuffer {
	cb := make([][]map[string]interface{}, width)
	for x := 0; x < width; x++ {
		cb[x] = make([]map[string]interface{}, height)
		for y := 0; y < height; y++ {
			cb[x][y] = make(map[string]interface{})
		}
	}
	return cb
}

// CanvasRenderer is used to render to HTML5 canvas in the browser when compiling to webassembly.
type CanvasRenderer struct {
	Buffer CanvasBuffer
	Width  int
	Height int
}

func NewScreen() Screen {
	bg := ColorFromPalette(Black, Black).Foreground
	js.Global().Call("initializeScreen", fmt.Sprintf("rgb(%d, %d, %d)", bg.R, bg.G, bg.B))
	size := js.Global().Call("size")
	w, h := size.Get("width").Int(), size.Get("height").Int()
	return &CanvasRenderer{Buffer: NewCanvasBuffer(w, h), Width: w, Height: h}
}

func (cr *CanvasRenderer) CleanUp() {

}

func (cr *CanvasRenderer) Clear() {
	cr.Buffer = NewCanvasBuffer(cr.Width, cr.Height)
	js.Global().Call("clear")
}

func (cr *CanvasRenderer) PollEvent() ScreenEvent {
	wait := make(chan js.Value)
	js.Global().Call("pollEvent").Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		wait <- args[0]
		close(wait)
		return nil
	}))
	jsEvent := <-wait
	switch jsEvent.Get("type").String() {
	case "key":
		switch jsEvent.Get("key").String() {
		case "ArrowLeft":
			return &KeyEvent{Key: KeyLeft}
		case "ArrowRight":
			return &KeyEvent{Key: KeyRight}
		case "ArrowUp":
			return &KeyEvent{Key: KeyUp}
		case "ArrowDown":
			return &KeyEvent{Key: KeyDown}
		}
	case "mouse":
		return &MouseEvent{}
	}
	return nil
}

func (cr *CanvasRenderer) PostEvent(ev ScreenEvent) error {
	switch event := ev.(type) {
	case *KeyEvent:
		js.Global().Call("postKeyEvent", event.Rune)
	case *MouseEvent:
		js.Global().Call("postMouseEvent", event.X, event.Y, event.Button)
	}
	return nil
}

func (cr *CanvasRenderer) SetCellContent(x int, y int, primary rune, style Color) {
	cell := map[string]interface{}{
		"text":       string(primary),
		"foreground": fmt.Sprintf("rgb(%d, %d, %d)", style.Foreground.R, style.Foreground.G, style.Foreground.B),
		"background": fmt.Sprintf("rgb(%d, %d, %d)", style.Background.R, style.Background.G, style.Background.B),
	}
	cr.Buffer[x][y] = cell
}

func (cr *CanvasRenderer) Show() {
	bytes, err := json.Marshal(cr.Buffer)
	if err != nil {
		js.Global().Get("console").Call("error", err)
	}
	js.Global().Call("show", string(bytes))
}

func (cr *CanvasRenderer) Size() (int, int) {
	val := js.Global().Call("size")
	return val.Get("width").Int(), val.Get("height").Int()
}

func (cr *CanvasRenderer) Sync() {
	size := js.Global().Call("size")
	cr.Width, cr.Height = size.Get("width").Int(), size.Get("height").Int()
}

var palette [len(Colors)]string

func init() {
	for i, color := range Colors {
		palette[i] = fmt.Sprint("#", color)
	}
}

// ColorFromPalette resolves color shortcuts by name
func ColorFromPalette(foreground, background ColorName) string {
	return Color{palette[foreground], palette[background]}
}
