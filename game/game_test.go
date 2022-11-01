package game_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game"
	"github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/events"
	"github.com/vgalaktionov/roguelike-go/game/resources"
	"github.com/vgalaktionov/roguelike-go/game/systems"
)

// RandomPlayer is a system used for testing.
// By registering this, it is possible to simulate a game without player input
func RandomPlayer(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)
	state := ecs.GetResource[resources.GameState](w)
	if state != resources.PlayerTurn {
		return
	}
	switch rand.Intn(4) {
	case 0:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyLeft, Rune: ' '})
	case 1:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyRight, Rune: ' '})
	case 2:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyUp, Rune: ' '})
	case 3:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyDown, Rune: ' '})
	}
}

func NewSimulatedGameRoomsAndCorridors() *game.Game {
	rand.Seed(time.Now().UnixNano())
	encoding.Register()

	renderer := tcell.NewSimulationScreen("UTF-8")
	renderer.Init()
	renderer.SetSize(200, 200)
	renderer.SetStyle(draw.DEFAULT_STYLE)
	renderer.Clear()
	screen := &draw.ConsoleRenderer{Screen: renderer}

	w := ecs.NewWorld()

	ecs.RegisterEventSystem(w, systems.Console, events.ConsoleEvent{})

	ecs.RegisterSystem(w, systems.UpdateVisibility)
	ecs.RegisterSystem(w, RandomPlayer)

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
			Style: draw.ColorFromPalette(draw.Yellow, draw.Black),
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
			components.Renderable{Glyph: glyph, Style: draw.ColorFromPalette(draw.Red, draw.Black)},
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
	return &game.Game{w, screen}
}

func BenchmarkRandomGameWithRoomsAndCorridors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game := NewSimulatedGameRoomsAndCorridors()
		game.RunN(100)
	}
}
