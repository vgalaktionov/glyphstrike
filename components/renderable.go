package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Renderable struct {
	Glyph rune
	Style tcell.Style
}

const RenderableTag = ecs.CTag("Renderable")

func (Renderable) CTag() ecs.CTag {
	return RenderableTag
}
