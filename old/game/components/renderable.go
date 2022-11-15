package components

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// Renderable is an entity that we know how to draw to the screen.
type Renderable struct {
	Glyph      rune
	Foreground draw.ColorName
	Background draw.ColorName
}

func (Renderable) CID() ecs.CID {
	return renderableID
}
