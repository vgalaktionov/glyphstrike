package components

import (
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Viewshed struct {
	Radius int
	View   *fov.View
}

const ViewshedTag = ecs.CTag("Viewshed")

func (Viewshed) CTag() ecs.CTag {
	return ViewshedTag
}
