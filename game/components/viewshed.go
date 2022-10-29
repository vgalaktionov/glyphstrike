package components

import (
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// Viewshed is a fancy name for "the things you can see", in our case for a given entity.
type Viewshed struct {
	Radius int
	View   *fov.View
}

const ViewshedTag = ecs.CTag("Viewshed")

func (Viewshed) CTag() ecs.CTag {
	return ViewshedTag
}
