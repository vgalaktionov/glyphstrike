package main

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/engine"
)

func main() {
	e := engine.NewEngine()

	log.SetFlags(0)
	log.SetOutput(e.Logger())

	quit := func() {
		maybePanic := recover()
		e.Renderer.Clear()
		e.Renderer.ShowCursor(0, 0)
		e.Renderer.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	e.Run()
}
