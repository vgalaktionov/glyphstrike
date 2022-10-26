package engine

import (
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Engine struct {
	ECS      *ecs.World
	Renderer draw.Renderer
}

func NewEngine() *Engine {
	rand.Seed(time.Now().UnixNano())
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

func (e *Engine) Run() {
	for {
		e.tick()
	}
}

func (e *Engine) tick() {
	e.Renderer.Clear()

	e.ECS.RunSystems(e.Renderer)

	e.Renderer.Show()
}
