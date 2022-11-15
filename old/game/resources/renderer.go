package resources

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Renderer struct {
	draw.Screen
}

func (Renderer) RID() ecs.RID {
	return rendererID
}
