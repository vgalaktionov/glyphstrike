package resources

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Renderer struct {
	*draw.ConsoleRenderer
}

const RendererTag = ecs.RTag("Renderer")

func (*Renderer) RTag() ecs.RTag {
	return RendererTag
}
