package main

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/game"
)

func main() {
	g := game.NewGame()

	log.SetFlags(0)
	log.SetOutput(g.Logger())

	quit := func() {
		maybePanic := recover()
		g.Renderer.Clear()
		g.Renderer.ShowCursor(0, 0)
		g.Renderer.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	g.Run()
}
