package systems

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

const UIOffsetX = 42
const UIOffsetY = 1

const title = "☄️ Glyphstrike"

func UI(r draw.Renderer, w *ecs.World) {
	maxX, maxY := r.Size()
	draw.DrawBox(r, 0, 0, UIOffsetX, maxY-1, tcell.StyleDefault, "")

	draw.DrawStr(r, maxX/2-len(title)/2, 0, tcell.StyleDefault, title)
}
