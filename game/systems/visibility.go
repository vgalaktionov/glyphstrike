package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// Update visibility computes visibility field for affected entities each tick using a raycasting algorithm.
func UpdateVisibility(w *ecs.World) {

	for e := range ecs.QueryEntitiesIter(w, Viewshed{}, Position{}) {
		viewshed := ecs.GetEntityComponent[Viewshed](w, e)
		pos := ecs.GetEntityComponent[Position](w, e)
		m := ecs.GetResource[*resources.Map](w)

		viewshed.View.Compute(m, pos.X, pos.Y, viewshed.Radius)
	}
}
