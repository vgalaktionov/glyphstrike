package draw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type Renderer interface {
	Sync()
	PollEvent() tcell.Event
	Clear()
	Show()
	Size() (height int, width int)
	SetContent(x, y int, primary rune, combining []rune, style tcell.Style)
	ShowCursor(x, y int)
	Fini()
}

func DrawStr(r Renderer, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		r.SetContent(x, y, c, comb, style)
		x += w
	}
}

func DrawBox(r Renderer, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			r.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		r.SetContent(col, y1, tcell.RuneHLine, nil, style)
		r.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		r.SetContent(x1, row, tcell.RuneVLine, nil, style)
		r.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		r.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		r.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		r.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		r.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	DrawStr(r, x1+1, y1+1, style, text)
}

var DEFAULT_STYLE tcell.Style = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
