package components

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// Renderable is an entity that we know how to draw to the screen.
type Renderable struct {
	Glyph rune
	Style draw.Color
}

func (Renderable) CTag() ecs.CTag {
	return ecs.CTag("Renderable")
}
