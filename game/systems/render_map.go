package systems

import (
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/components"

	//lint:ignore ST1001 dot importing resources makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/resources"
)

// RenderMap system is responsible for rendering the map, taking into account player visibility.
func RenderMap(w *ecs.World) {
	m := ecs.GetResource[*Map](w)
	r := ecs.GetResource[Renderer](w)

	e := ecs.GetResource[PlayerEntity](w).Entity()
	viewshed := ecs.MustGetEntityComponent[Viewshed](w, e)

	for x, col := range m.Tiles {
		renderX := x + UIOffsetX

	inner:
		for y, tile := range col {
			renderY := y + UIOffsetY
			if !viewshed.View.IsVisible(x, y) {
				// Display revealed tiles as greyed out
				if m.RevealedTiles[x][y] {
					switch tile {
					case FloorTile:
						r.SetCellContent(renderX, renderY, ' ', draw.DarkGray, draw.DarkGray)
					case WallTile:
						r.SetCellContent(renderX, renderY, '█', draw.DarkerGray, draw.Black)
					}
				}
				// If tile is neither visible nor revealed, skip rendering
				continue inner
			}

			// Keep track of tiles we reveal
			m.RevealedTiles[x][y] = true

			switch tile {
			case FloorTile:
				r.SetCellContent(renderX, renderY, ' ', draw.LightGray, draw.Gray)
			case WallTile:
				r.SetCellContent(renderX, renderY, '█', draw.BlueGreen, draw.Black)
			}
		}
	}
}
