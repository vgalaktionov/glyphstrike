package engine

import (
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/systems"
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
	screen.Clear()

	world := ecs.NewWorld()
	return &Engine{world, screen}
}

func (e *Engine) Run() {
	e.ECS.RunEventSystems(e.Renderer)

	for {
		e.tick()
	}
}

func (e *Engine) tick() {
	systems.ClearMap(e.Renderer)
	e.ECS.RunSystems(e.Renderer)
	e.Renderer.Show()
}
