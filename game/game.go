package game

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2/encoding"
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/events"
	"github.com/vgalaktionov/roguelike-go/game/resources"
	"github.com/vgalaktionov/roguelike-go/game/systems"
)

// Game binds together all the high-level subsystems for the game to use.
// Currently, this only means the ECS and the renderer.
type Game struct {
	ECS      *ecs.World
	Renderer draw.Screen
}

// NewGame abstracts away the technical details of setting up a terminal to render to.
// It also sets up default systems and events.
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	encoding.Register()

	screen := draw.NewScreen()

	w := ecs.NewWorld()

	ecs.RegisterEventSystem(w, systems.Console, events.ConsoleEvent{})

	ecs.RegisterSystem(w, systems.UpdateVisibility)

	// only one of these two will run
	ecs.RegisterSystem(w, systems.HandlePlayerInput)
	ecs.RegisterSystem(w, systems.ProcessMonsterAI)

	ecs.RegisterSystem(w, systems.MapIndexing)
	ecs.RegisterSystem(w, systems.MeleeCombat)
	ecs.RegisterSystem(w, systems.ApplyDamage)
	ecs.RegisterSystem(w, systems.Reap)

	ecs.RegisterSystem(w, systems.RenderMap)
	ecs.RegisterSystem(w, systems.Render)
	ecs.RegisterSystem(w, systems.UI)

	ecs.RegisterSystem(w, systems.UpdateTurn)

	screenX, screenY := screen.Size()
	mapX := screenX - systems.UIOffsetX
	mapY := screenY - systems.UIOffsetY

	m := resources.NewMapRoomsAndCorridors(mapX, mapY)

	ecs.SetResource(w, m)
	ecs.SetResource(w, resources.Renderer{Screen: screen})
	ecs.SetResource(w, resources.PreRun)

	playerX, playerY := m.Rooms[0].Center()
	ecs.AddEntity(
		w,
		components.Player{},
		components.Position{X: playerX, Y: playerY},
		components.Renderable{
			Glyph: '@',
			Style: draw.ColorFromPalette(draw.Yellow, draw.Transparent),
		},
		components.Viewshed{Radius: 8, View: fov.New()},
		components.Name("Player"),
		components.CombatStats{
			MaxHP:   30,
			HP:      30,
			Defense: 2,
			Power:   5,
		},
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

		ecs.AddEntity(
			w,
			components.Position{X: x, Y: y},
			components.Renderable{Glyph: glyph, Style: draw.ColorFromPalette(draw.Red, draw.Transparent)},
			components.Viewshed{Radius: 8, View: fov.New()},
			components.MonsterAI{},
			components.Name(fmt.Sprintf("Monster #%d", i)),
			components.BlocksTile{},
			components.CombatStats{
				MaxHP:   16,
				HP:      16,
				Defense: 1,
				Power:   4,
			},
		)
	}
	return &Game{w, screen}
}

// GameLogger retains a private reference to the Game
type GameLogger struct {
	game *Game
}

// Write implements the io.Writer interface for the console logging system
func (gl *GameLogger) Write(msg []byte) (int, error) {
	ecs.DispatchEvent(gl.game.ECS, events.ConsoleEvent{Message: string(msg)})
	return len(msg), nil
}

// Logger returns a log sink hooked up to the console system.
func (g *Game) Logger() io.Writer {
	return &GameLogger{game: g}
}

// CleanUp performs any necessary cleanup actions on crash or exit.
func (g *Game) CleanUp() {
	g.Renderer.CleanUp()
}

// Run is the entrypoint into the Game. It kicks off background (event) systems and starts the game loop.
func (g *Game) Run() {
	ecs.RunEventSystems(g.ECS)

	for {
		g.tick()
	}
}

// tick is called on every tick of the game loop. It clears the map portion of the screen, processes
// all blocking systems and flushes the terminal buffer to screen.
func (g *Game) tick() {
	ecs.RunSystems(g.ECS)
	g.Renderer.Show()
}
