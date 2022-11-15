package systems

import (
	"fmt"
	"math"

	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

const UIOffsetX = 42
const UIOffsetY = 1
const ConsoleOffsetY = 20 + UIOffsetY

// UI system render all elements besides map and console to the screen.
// It runs as a normal blocking system and updates once per tick.
func UI(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)
	playerEnt := ecs.GetResource[resources.PlayerEntity](w).Entity()
	_, maxY := r.Size()

	// clear UI area except for console
	draw.Fill(r, 1, UIOffsetY+1, UIOffsetX, ConsoleOffsetY, draw.Black)
	// draw outer border
	draw.DrawBox(r, 0, UIOffsetY, UIOffsetX, maxY-1, draw.White, draw.Black, "")

	stats := ecs.MustGetEntityComponent[components.CombatStats](w, playerEnt)
	renderHP(r, 1, UIOffsetY+1, stats)

	renderConsole(r, 0, ConsoleOffsetY, UIOffsetX, maxY-1)

	renderTooltips(w)
}

func renderHP(r draw.Screen, x, y int, stats components.CombatStats) {
	barMaxWidth := UIOffsetX - 2

	draw.DrawStr(r, x, y, draw.White, draw.Black, fmt.Sprintf("HP: %d/%d", stats.HP, stats.MaxHP))
	barWidth := int(math.Floor((float64(stats.HP) / float64(stats.MaxHP)) * float64(barMaxWidth)))

	draw.DrawHBar(r, x, x+barWidth, y+1, draw.Red)
}

func renderConsole(r draw.Screen, x1, y1, x2, y2 int) {
	draw.DrawBox(r, x1, y1, x2, y2, draw.White, draw.Black, "")
	draw.DrawStr(r, x1+1, y1, draw.White, draw.Black, "Console")
}

func renderTooltips(w *ecs.World) {
	mousePos := ecs.GetResource[resources.MousePosition](w)
	r := ecs.GetResource[resources.Renderer](w)

	if mousePos.Moved {
		r.SetCellContent(mousePos.X, mousePos.Y, ' ', draw.Red, draw.Red)
		ecs.SetResource(w, resources.MousePosition{X: mousePos.X, Y: mousePos.Y})
	}
}
