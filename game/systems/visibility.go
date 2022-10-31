package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// Update visibility computes visibility field for affected entities each tick using a raycasting algorithm.
func UpdateVisibility(w *ecs.World) {

	for e := range ecs.QueryEntitiesIter(w, Viewshed{}, Position{}) {
		viewshed := ecs.MustGetEntityComponent[Viewshed](w, e)
		pos := ecs.MustGetEntityComponent[Position](w, e)
		m := ecs.GetResource[*resources.Map](w)

		viewshed.View.Compute(m, pos.X, pos.Y, viewshed.Radius)
	}
}
