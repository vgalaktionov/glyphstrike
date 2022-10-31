//go:build js
// +build js

package draw

import "github.com/gdamore/tcell/v2"

// CanvasRender is used to render to HTML5 canvas in the browser when compiling to webassembly.
type CanvasRenderer struct {
}

func NewScreen() Screen {
	return &CanvasRenderer{}
}

func (cr *CanvasRenderer) CleanUp() {

}

func (cr *CanvasRenderer) Clear() {

}

func (cr *CanvasRenderer) PollEvent() tcell.Event {

}

func (cr *CanvasRenderer) PostEvent(ev tcell.Event) error {

}

func (cr *CanvasRenderer) SetContent(x int, y int, primary rune, combining []rune, style tcell.Style) {

}

func (cr *CanvasRenderer) Show() {

}

func (cr *CanvasRenderer) Size() (int, int) {

}

func (cr *CanvasRenderer) Sync() {

}
