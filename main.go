package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/engine"
	"github.com/vgalaktionov/roguelike-go/systems"
)

func main() {
	e := engine.NewEngine()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		e.Screen.Clear()
		e.Screen.ShowCursor(0, 0)
		e.Screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// State
	e.ECS.RegisterSystem(&systems.PlayerInput{})
	e.ECS.RegisterSystem(&systems.LeftWalker{})
	e.ECS.RegisterSystem(&systems.Render{})

	e.ECS.AddEntity(
		components.Player{},
		components.Position{X: 40, Y: 25},
		components.Renderable{
			Glyph: '@',
			Style: tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorBlack),
		},
	)

	for i := 0; i < 10; i++ {
		e.ECS.AddEntity(
			components.LeftMover{},
			components.Position{X: i * 7, Y: 20},
			components.Renderable{
				Glyph: '☺',
				Style: tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack),
			},
		)

	}
	// Event loop
	for {
		e.Tick()
	}
}
