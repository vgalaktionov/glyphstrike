package engine

import (
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/events"
	"github.com/vgalaktionov/roguelike-go/systems"
)

// Engine binds together all the high-level subsystems for the game to use.
// Currently, this only means the ECS and the renderer.
type Engine struct {
	ECS      *ecs.World
	Renderer draw.Renderer
}

// NewEngine abstracts away the technical details of setting up a terminal to render to.
// It also sets up default systems and events.
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

	world.RegisterSystem(systems.HandlePlayerInput)
	world.RegisterSystem(systems.UpdateVisibility)
	world.RegisterSystem(systems.ProcessMonsterAI)
	world.RegisterEventSystem(systems.Console)
	world.RegisterEvent(events.ConsoleEvent{})
	world.RegisterSystem(systems.RenderMap)
	world.RegisterSystem(systems.Render)
	world.RegisterSystem(systems.UI)

	screenX, screenY := screen.Size()
	mapX := screenX - systems.UIOffsetX
	mapY := screenY - systems.UIOffsetY

	m := systems.NewMapRoomsAndCorridors(mapX, mapY)
	// m := systems.NewTestMap(mapX, mapY)
	world.AddResource(m)

	playerX, playerY := m.Rooms[0].Center()
	world.AddEntity(
		components.Player{},
		components.Position{X: playerX, Y: playerY},
		components.Renderable{
			Glyph: '⊛',
			Style: tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack.TrueColor()),
		},
		components.Viewshed{Radius: 8, View: fov.New()},
	)

	for i := 1; i < len(m.Rooms); i++ {
		x, y := m.Rooms[i].Center()

		var glyph rune
		roll := rand.Intn(2)
		switch roll {
		case 0:
			glyph = 'g' // goblin
		case 1:
			glyph = 'o' // orc
		}

		world.AddEntity(
			components.Position{X: x, Y: y},
			components.Renderable{Glyph: glyph, Style: tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor())},
			components.Viewshed{Radius: 8, View: fov.New()},
			components.MonsterAI{},
		)
	}
	return &Engine{world, screen}
}

// Run is the entrypoint into the engine. It kicks off background (event) systems and starts the game loop.
func (e *Engine) Run() {
	e.ECS.RunEventSystems(e.Renderer)

	for {
		e.tick()
	}
}

// tick is called on every tick of the game loop. It clears the map portion of the screen, processes
// all blocking systems and flushes the terminal buffer to screen.
func (e *Engine) tick() {
	systems.ClearMap(e.Renderer)
	e.ECS.RunSystems(e.Renderer)
	e.Renderer.Show()
}
