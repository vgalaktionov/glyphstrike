package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/vgalaktionov/roguelike-go/events"
	"github.com/vgalaktionov/roguelike-go/render"
	"github.com/vgalaktionov/roguelike-go/state"
)

func main() {
	encoding.Register()
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(render.DEFAULT_STYLE)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Clear()
		s.ShowCursor(0, 0)
		s.Fini()
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
	w := state.NewWorld(s.Size())
	// Event loop
	for {
		render.Draw(s, w)
		events.HandleEvents(s, w)
	}
}
