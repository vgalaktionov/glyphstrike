package systems

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/components"

	//lint:ignore ST1001 dot importing resources makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/resources"
)

// ClearMap only clears the map part of the screen, leaving UI elements intact.
func clearMap(r draw.Screen) {
	maxX, maxY := r.Size()
	for x := UIOffsetX; x <= maxX; x++ {
		for y := UIOffsetY; y <= maxY; y++ {
			r.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}
}

// RenderMap system is responsible for rendering the map, taking into account player visibility.
func RenderMap(w *ecs.World) {
	m := ecs.GetResource[*Map](w)
	r := ecs.GetResource[Renderer](w)

	clearMap(r)

	for e := range ecs.QueryEntitiesIter(w, Player{}, Viewshed{}) {
		viewshed := ecs.GetEntityComponent[Viewshed](w, e)

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
							r.SetContent(renderX, renderY, ' ', nil, tcell.StyleDefault.Background(tcell.NewRGBColor(20, 20, 20)))
						case WallTile:
							r.SetContent(renderX, renderY, '█', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(50, 50, 50).TrueColor()))
						}
					}
					// If tile is neither visible nor revealed, skip rendering
					continue inner
				}

				// Keep track of tiles we reveal
				m.RevealedTiles[x][y] = true

				switch tile {
				case FloorTile:
					r.SetContent(renderX, renderY, ' ', nil, tcell.StyleDefault.Foreground(tcell.ColorLightGray.TrueColor()).Bold(true))
				case WallTile:
					r.SetContent(renderX, renderY, '█', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGreen.TrueColor()).Bold(true))
				}
			}
		}
	}
}
