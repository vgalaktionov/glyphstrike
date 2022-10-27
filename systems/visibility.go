package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/resources"
)

// Update visibility computes visibility field for affected entities each tick using a raycasting algorithm.
func UpdateVisibility(w *ecs.World) {

	for e := range w.QueryEntitiesIter(Viewshed{}, Position{}) {
		viewshed := w.GetEntityComponent(ViewshedTag, e).(Viewshed)
		pos := w.GetEntityComponent(PositionTag, e).(Position)
		m := w.GetResource(resources.MapTag).(*resources.Map)

		viewshed.View.Compute(m, pos.X, pos.Y, viewshed.Radius)
	}
}
