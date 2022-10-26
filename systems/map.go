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
	m := w.GetResource(MapTag).(*Map)
	for x, col := range m.Tiles {
		for y, tile := range col {
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

func NewEmptyMap(mapWidth, mapHeight int) *Map {
	m := &Map{Tiles: make([][]TileType, mapWidth), Width: mapWidth, Height: mapHeight}
	for i := 0; i < mapWidth; i++ {
		m.Tiles[i] = make([]TileType, m.Height)
	}
	return m
}

func NewTestMap(mapWidth, mapHeight, playerX, playerY int) *Map {
	m := NewEmptyMap(mapWidth, mapHeight)

	// Close off the sides with walls
	for x := 0; x < mapWidth; x++ {
		m.Tiles[x][0] = WallTile
		m.Tiles[x][m.Height-1] = WallTile
	}

	for y := 0; y < mapHeight; y++ {
		m.Tiles[0][y] = WallTile
		m.Tiles[m.Width-1][y] = WallTile
	}

	// Random walls, avoiding the player
	for i := 0; i < 800; i++ {
		x := rand.Intn(m.Width - 1)
		y := rand.Intn(m.Height - 1)
		if !(x == playerX && y == playerY) {
			m.Tiles[x][y] = WallTile
		}
	}

	return m
}

func NewMapRoomsAndCorridors(mapWidth, mapHeight, playerX, playerY int) *Map {
	m := NewEmptyMap(mapWidth, mapHeight)
	m.Fill(WallTile)

	room1 := draw.NewRect(20, 15, 10, 15)
	room2 := draw.NewRect(65, 5, 20, 25)

	m.ApplyRoom(room1)
	m.ApplyRoom(room2)
	m.ApplyHorizontalTunnel(25, 75, 20)

	return m
}
