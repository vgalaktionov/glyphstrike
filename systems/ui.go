package systems

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

const UIOffsetX = 32
const UIOffsetY = 3

func UI(r draw.Renderer, w *ecs.World) {
	_, maxY := r.Size()
	draw.DrawBox(r, 0, 0, UIOffsetX, maxY-1, tcell.StyleDefault, "")
}
