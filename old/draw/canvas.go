//go:build js
// +build js

package draw

import (
	"bytes"
	"fmt"
	"syscall/js"
)

// CanvasBufferCell is a single cell of the canvasbuffer
type CanvasBufferCell struct {
	Foreground, Background ColorName
	Char                   rune
}

// CanvasBuffer is the internal buffer of cells to be drawn with the next CanvasRenderer.Show() call.
type CanvasBuffer [][]CanvasBufferCell

func NewCanvasBuffer(width, height int) CanvasBuffer {
	cb := make([][]CanvasBufferCell, width)
	for x := 0; x < width; x++ {
		cb[x] = make([]CanvasBufferCell, height)
		for y := 0; y < height; y++ {
			cb[x][y] = CanvasBufferCell{Black, Black, ' '}
		}
	}
	return cb
}

// CanvasRenderer is used to render to HTML5 canvas in the browser when compiling to webassembly.
type CanvasRenderer struct {
	Renderer js.Value
	Buffer   CanvasBuffer
	Bytes    *bytes.Buffer
	Width    int
	Height   int
}

const LineLength = 4 + 6 + 6 // max 4 byte rune + 6 byte foreground hex + 6 byte background hex

func NewScreen() Screen {
	renderer := js.Global().Get("CanvasRenderer").New(fmt.Sprint("#", palette[Black]))
	js.Global().Set("renderer", renderer)
	size := renderer.Call("size")
	w, h := size.Get("width").Int(), size.Get("height").Int()
	buf := new(bytes.Buffer)
	return &CanvasRenderer{Renderer: renderer, Buffer: NewCanvasBuffer(w, h), Bytes: buf, Width: w, Height: h}
}

// NewSimulatedScreen is just a screen for canvas
func NewSimulatedScreen() Screen {
	return NewScreen()
}

func (cr *CanvasRenderer) CleanUp() {

}

func (cr *CanvasRenderer) Clear() {
	cr.Buffer = NewCanvasBuffer(cr.Width, cr.Height)
	cr.Renderer.Call("clear")
}

func (cr *CanvasRenderer) PollEvent() ScreenEvent {
	wait := make(chan js.Value)
	cr.Renderer.Call("pollEvent").Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
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
		default:
			return &KeyEvent{Key: KeyRune, Rune: rune(jsEvent.Get("key").String()[0])}
		}
	case "mouse":
		btn := jsEvent.Get("button").Int()
		var button MouseButton
		if btn == 0 {
			button = Primary
		} else if btn == 1 {
			button = Secondary
		}
		return &MouseEvent{X: jsEvent.Get("x").Int(), Y: jsEvent.Get("y").Int(), Button: button}
	}
	return nil
}

func (cr *CanvasRenderer) PostEvent(ev ScreenEvent) error {
	switch event := ev.(type) {
	case *KeyEvent:
		cr.Renderer.Call("postKeyEvent", event.Rune)
	case *MouseEvent:
		cr.Renderer.Call("postMouseEvent", event.X, event.Y, event.Button)
	}
	return nil
}

func (cr *CanvasRenderer) SetCellContent(x int, y int, primary rune, foreground, background ColorName) {
	cr.Buffer[x][y] = CanvasBufferCell{foreground, background, primary}
}

func (cr *CanvasRenderer) Show() {
	cr.Bytes.Reset()

	go func() {
		for _, column := range cr.Buffer {
			for _, cell := range column {
				for padding := 4 - len([]byte(string(cell.Char))); padding > 0; padding-- {
					cr.Bytes.WriteRune(' ')
				}
				cr.Bytes.WriteRune(cell.Char)
				cr.Bytes.WriteString(palette[cell.Foreground])
				cr.Bytes.WriteString(palette[cell.Background])
			}
		}
		jsBuffer := js.Global().Get("Uint8Array").New(cr.Width * cr.Height * LineLength)
		js.CopyBytesToJS(jsBuffer, cr.Bytes.Bytes())

		cr.Renderer.Call("show", jsBuffer)
	}()
}

func (cr *CanvasRenderer) Size() (int, int) {
	val := cr.Renderer.Call("size")
	return val.Get("width").Int(), val.Get("height").Int()
}

func (cr *CanvasRenderer) Sync() {
	cr.Renderer.Call("sync")
	size := cr.Renderer.Call("size")
	cr.Width, cr.Height = size.Get("width").Int(), size.Get("height").Int()
}

var palette [len(Colors)]string

func init() {
	for i, color := range Colors {
		palette[i] = fmt.Sprintf("%x", color)
	}
}
