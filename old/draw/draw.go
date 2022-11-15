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
	SetCellContent(x, y int, primary rune, foreground, background ColorName)
	CleanUp()
}

// DrawStr draws a single line of text to the screen. This will not work correctly for combining characters.
func DrawStr(r Screen, x, y int, foreground, background ColorName, str string) {
	for _, c := range str {
		r.SetCellContent(x, y, c, foreground, background)
		x += 1
	}
}

// DrawHBar draws a single colored horizontal bar to the screen.
func DrawHBar(r Screen, x1, x2, y int, color ColorName) {
	for x := x1; x < x2; x++ {
		r.SetCellContent(x, y, '█', color, color)
	}
}

// DrawBox draws a rectangular box with the standard box drawing characters, and (optional) text contents.
func DrawBox(r Screen, x1, y1, x2, y2 int, foreground, background ColorName, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		r.SetCellContent(col, y1, tcell.RuneHLine, foreground, background)
		r.SetCellContent(col, y2, tcell.RuneHLine, foreground, background)
	}
	for row := y1 + 1; row < y2; row++ {
		r.SetCellContent(x1, row, tcell.RuneVLine, foreground, background)
		r.SetCellContent(x2, row, tcell.RuneVLine, foreground, background)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		r.SetCellContent(x1, y1, tcell.RuneULCorner, foreground, background)
		r.SetCellContent(x2, y1, tcell.RuneURCorner, foreground, background)
		r.SetCellContent(x1, y2, tcell.RuneLLCorner, foreground, background)
		r.SetCellContent(x2, y2, tcell.RuneLRCorner, foreground, background)
	}

	DrawStr(r, x1+1, y1+1, foreground, background, text)
}

// FIll fills a rectangular box with the provided color.
func Fill(r Screen, x1, y1, x2, y2 int, color ColorName) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		for row := y1; row < y2; row++ {
			r.SetCellContent(col, row, ' ', color, color)
		}
	}
}
