package engine

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Engine struct {
	ECS    *ecs.World
	Screen tcell.Screen
}

func NewEngine() *Engine {
	encoding.Register()
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(draw.DEFAULT_STYLE)
	screen.EnableMouse()
	screen.EnablePaste()
	screen.Clear()

	world := ecs.NewWorld()
	return &Engine{world, screen}
}

func (e *Engine) Tick() {
	e.Screen.Clear()
	e.ECS.Process(e.Screen)
	e.Screen.Show()
}
