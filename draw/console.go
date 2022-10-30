package draw

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type ConsoleRenderer struct {
	tcell.Screen
}

func NewConsoleRenderer() *ConsoleRenderer {
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(DEFAULT_STYLE)
	screen.EnableMouse()
	screen.Clear()
	return &ConsoleRenderer{screen}
}

func (cr *ConsoleRenderer) CleanUp() {
	cr.Clear()
	cr.ShowCursor(0, 0)
	cr.Fini()
}
