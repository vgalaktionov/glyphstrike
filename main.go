package main

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/game"
)

func main() {
	g := game.NewGame()

	// Hook up stdlib logging output to the game window console
	log.SetFlags(0)
	log.SetOutput(g.Logger())

	quit := func() {
		maybePanic := recover()
		g.CleanUp()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	g.Run()
}
