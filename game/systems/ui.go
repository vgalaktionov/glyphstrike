package systems

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

const UIOffsetX = 32
const UIOffsetY = 1

// UI system render all elements besides map and console to the screen.
// It runs as a normal blocking system and updates once per tick.
func UI(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)
	_, maxY := r.Size()
	// draw console
	draw.DrawBox(r, 0, UIOffsetY, UIOffsetX, maxY-1, draw.White, draw.Black, "")
}
