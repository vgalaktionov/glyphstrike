package systems

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"

	//lint:ignore ST1001 dot importing resources makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/resources"
)

func RenderMap(r draw.Renderer, w *ecs.World) {
	m := w.GetResource(MapTag).(Map)
	for x, row := range m {
		for y, tile := range row {
			// render tiles
			switch tile {
			case FloorTile:
				r.SetContent(x, y, '.', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGrey))
			case WallTile:
				r.SetContent(x, y, '#', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGreen))
			}
		}
	}
}

func NewMap(mapWidth, mapHeight, playerX, playerY int) Map {
	m := make(Map, mapWidth)
	for i := 0; i < mapWidth; i++ {
		m[i] = make([]TileType, mapHeight)
	}

	// Close off the sides with walls
	for x := 0; x < mapWidth; x++ {
		m[x][0] = WallTile
		m[x][mapHeight-1] = WallTile
	}

	for y := 0; y < mapHeight; y++ {
		m[0][y] = WallTile
		m[mapWidth-1][y] = WallTile
	}

	// Random walls, avoiding the player
	for i := 0; i < 800; i++ {
		x := rand.Intn(mapWidth - 1)
		y := rand.Intn(mapHeight - 1)
		if !(x == playerX && y == playerY) {
			m[x][y] = WallTile
		}
	}

	return m
}
