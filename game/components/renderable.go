package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// Renderable is an entity that we know how to draw to the screen.
type Renderable struct {
	Glyph rune
	Style tcell.Style
}

func (Renderable) CTag() ecs.CTag {
	return ecs.CTag("Renderable")
}