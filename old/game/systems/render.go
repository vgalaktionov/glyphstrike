package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// Render system handles drawing non-map renderable entities to the screen, taking visibility into account.
func Render(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)

	playerEnt := ecs.GetResource[resources.PlayerEntity](w).Entity()
	playerViewshed := ecs.MustGetEntityComponent[Viewshed](w, playerEnt)

	for _, e := range ecs.QueryEntitiesIter(w, Renderable{}, Position{}) {
		pos := ecs.MustGetEntityComponent[Position](w, e)

		if !playerViewshed.View.IsVisible(pos.X, pos.Y) {
			continue
		}

		renderable := ecs.MustGetEntityComponent[Renderable](w, e)
		r.SetCellContent(pos.X+UIOffsetX, pos.Y+UIOffsetY, renderable.Glyph, renderable.Foreground, renderable.Background)
	}
}
