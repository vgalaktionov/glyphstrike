package events

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/state"
)

type Action interface {
	Handle(w *state.World)
}

func HandleEvents(s tcell.Screen, world *state.World) {
	// Poll event
	ev := s.PollEvent()

	// Process event
	switch ev := ev.(type) {
	case *tcell.EventResize:
		s.Sync()

	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			s.Clear()
			s.ShowCursor(0, 0)
			s.Fini()
			os.Exit(0)
		}

		var action Action
		switch ev.Key() {
		case tcell.KeyLeft:
			action = &MovementAction{-1, 0}
		case tcell.KeyRight:
			action = &MovementAction{1, 0}
		case tcell.KeyUp:
			action = &MovementAction{0, -1}
		case tcell.KeyDown:
			action = &MovementAction{0, 1}
		}
		action.Handle(world)

	case *tcell.EventMouse:
		// mouseX, mouseY := ev.Position()

		switch ev.Buttons() {
		case tcell.Button1, tcell.Button2:
		case tcell.ButtonNone:
		}
	}
}
