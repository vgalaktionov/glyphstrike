package draw

import (
	"github.com/gdamore/tcell/v2"
)

// Screen abstracts away the renderer (currently only tcell), to ease testing as well as
// implementation of different rendering backends in the future.
type Screen interface {
	Sync()
	PollEvent() ScreenEvent
	PostEvent(ev ScreenEvent) error
	Clear()
	Show()
	Size() (height int, width int)
	SetCellContent(x, y int, primary rune, color Color)
	CleanUp()
}

// DrawStr draws a single line of text to the screen. This will not work correctly for combining characters.
func DrawStr(r Screen, x, y int, color Color, str string) {
	for _, c := range str {
		r.SetCellContent(x, y, c, color)
		x += 1
	}
}

// DrawBox draws a rectangular box with the standard box drawing characters, and (optional) text contents.
func DrawBox(r Screen, x1, y1, x2, y2 int, color Color, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		r.SetCellContent(col, y1, tcell.RuneHLine, color)
		r.SetCellContent(col, y2, tcell.RuneHLine, color)
	}
	for row := y1 + 1; row < y2; row++ {
		r.SetCellContent(x1, row, tcell.RuneVLine, color)
		r.SetCellContent(x2, row, tcell.RuneVLine, color)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		r.SetCellContent(x1, y1, tcell.RuneULCorner, color)
		r.SetCellContent(x2, y1, tcell.RuneURCorner, color)
		r.SetCellContent(x1, y2, tcell.RuneLLCorner, color)
		r.SetCellContent(x2, y2, tcell.RuneLRCorner, color)
	}

	DrawStr(r, x1+1, y1+1, color, text)
}

var DEFAULT_STYLE tcell.Style = tcell.StyleDefault.Background(tcell.ColorBlack.TrueColor()).Foreground(tcell.ColorWhite.TrueColor())
