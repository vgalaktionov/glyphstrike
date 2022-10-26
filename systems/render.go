package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

func Render(r draw.Renderer, w *ecs.World) {

	for e := range w.QueryEntitiesIter(Renderable{}, Position{}) {
		pos := w.GetEntityComponent(PositionTag, e).(Position)
		renderable := w.GetEntityComponent(RenderableTag, e).(Renderable)

		r.SetContent(pos.X, pos.Y, renderable.Glyph, nil, renderable.Style)
	}
}
