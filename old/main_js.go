//go:build js
// +build js

package main

import (
	"log"
	"syscall/js"

	"github.com/vgalaktionov/roguelike-go/game"
)

func main() {
	g := game.NewGame(false)

	// Hook up stdlib logging output to the game window console
	log.SetFlags(0)
	log.SetOutput(g.Logger())

	quit := func() {
		maybePanic := recover()
		g.CleanUp()
		if maybePanic != nil {
			js.Global().Get("console").Call("error", maybePanic)
			panic(maybePanic)
		}
	}
	defer quit()

	g.Run()
	<-make(chan struct{})
}
