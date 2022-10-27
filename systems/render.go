package systems

import (
	"fmt"

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// Render system handles drawing non-map renderable entities to the screen, taking visibility into account.
func Render(r draw.Renderer, w *ecs.World) {
	playerEnt := w.QueryEntitiesSingle(Player{})
	if playerEnt == ecs.EntityNotFound {
		panic(fmt.Sprintf("+%v", w))
	}
	playerViewshed := w.GetEntityComponent(ViewshedTag, playerEnt).(Viewshed)

	for e := range w.QueryEntitiesIter(Renderable{}, Position{}) {
		pos := w.GetEntityComponent(PositionTag, e).(Position)

		if !playerViewshed.View.IsVisible(pos.X, pos.Y) {
			continue
		}

		renderable := w.GetEntityComponent(RenderableTag, e).(Renderable)
		r.SetContent(pos.X+UIOffsetX, pos.Y+UIOffsetY, renderable.Glyph, nil, renderable.Style)
	}
}
