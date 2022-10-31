package systems

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

const UIOffsetX = 42
const UIOffsetY = 1

const title = "☄️ Glyphstrike"

// UI system render all elements besides map and console to the screen.
// It runs as a normal blocking system and updates once per tick.
func UI(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)
	maxX, maxY := r.Size()
	draw.DrawBox(r, 0, 0, UIOffsetX, maxY-1, tcell.StyleDefault, "")

	draw.DrawStr(r, maxX/2-len(title)/2, 0, tcell.StyleDefault, title)
}