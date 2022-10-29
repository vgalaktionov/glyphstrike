package systems

import (
	"fmt"

	"github.com/vgalaktionov/roguelike-go/ecs"
	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// Render system handles drawing non-map renderable entities to the screen, taking visibility into account.
func Render(w *ecs.World) {
	r := ecs.GetResource[*resources.Renderer](w)

	playerEnt := ecs.QueryEntitiesSingle(w, Player{})
	if playerEnt == ecs.EntityNotFound {
		panic(fmt.Sprintf("+%v", w))
	}
	playerViewshed := ecs.GetEntityComponent[Viewshed](w, playerEnt)

	for e := range ecs.QueryEntitiesIter(w, Renderable{}, Position{}) {
		pos := ecs.GetEntityComponent[Position](w, e)

		if !playerViewshed.View.IsVisible(pos.X, pos.Y) {
			continue
		}

		renderable := ecs.GetEntityComponent[Renderable](w, e)
		r.SetContent(pos.X+UIOffsetX, pos.Y+UIOffsetY, renderable.Glyph, nil, renderable.Style)
	}
}
