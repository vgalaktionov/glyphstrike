package main

import (
	"log"
	"os"

	"github.com/vgalaktionov/roguelike-go/engine"
)

func main() {
	f, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	e := engine.NewEngine()
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
