package systems

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"

	//lint:ignore ST1001 dot importing resources makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/resources"
)

func RenderMap(r draw.Renderer, w *ecs.World) {
	m := w.GetResource(MapTag).(*Map)
	for e := range w.QueryEntitiesIter(Player{}, Viewshed{}) {
		viewshed := w.GetEntityComponent(ViewshedTag, e).(Viewshed)
		for x, col := range m.Tiles {
		inner:
			for y, tile := range col {

				if !viewshed.View.IsVisible(x, y) {
					if m.RevealedTiles[x][y] {
						switch tile {
						case FloorTile:
							r.SetContent(x, y, '.', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGray.TrueColor()))
						case WallTile:
							r.SetContent(x, y, '#', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGray.TrueColor()))
						}
					}
					continue inner
				}

				m.RevealedTiles[x][y] = true

				switch tile {
				case FloorTile:
					r.SetContent(x, y, '.', nil, tcell.StyleDefault.Foreground(tcell.ColorLightGray.TrueColor()).Bold(true))
				case WallTile:
					r.SetContent(x, y, '#', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGreen.TrueColor()).Bold(true))
				}
			}
		}
	}
}

func NewEmptyMap(mapWidth, mapHeight int) *Map {
	m := &Map{Tiles: make([][]TileType, mapWidth), Width: mapWidth, Height: mapHeight, RevealedTiles: make([][]bool, mapWidth)}
	for i := 0; i < mapWidth; i++ {
		m.Tiles[i] = make([]TileType, m.Height)
		m.RevealedTiles[i] = make([]bool, m.Height)
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

func NewMapRoomsAndCorridors(mapWidth, mapHeight int) *Map {
	m := NewEmptyMap(mapWidth, mapHeight)
	m.Fill(WallTile)

	maxRooms := 30
	minSize := 6
	maxSize := 20

	for i := 0; i < maxRooms; i++ {
		w := rand.Intn(maxSize-minSize) + minSize
		h := rand.Intn(maxSize-minSize) + minSize
		x := rand.Intn(m.Width - w - 1)
		y := rand.Intn(m.Height - h - 1)

		room := draw.NewRect(x, y, w, h)
		ok := true
		for _, otherRoom := range m.Rooms {
			if room.Intersect(otherRoom) {
				ok = false
				break
			}
		}
		if ok {
			m.ApplyRoom(room)

			if len(m.Rooms) > 0 {
				newX, newY := room.Center()
				prevX, prevY := m.Rooms[len(m.Rooms)-1].Center()
				if rand.Intn(2) == 1 {
					m.ApplyHorizontalTunnel(prevX, newY, prevY)
					m.ApplyVerticalTunnel(prevY, newY, newX)
				} else {
					m.ApplyVerticalTunnel(prevY, newY, prevX)
					m.ApplyHorizontalTunnel(prevX, newX, newY)
				}
			}
			m.Rooms = append(m.Rooms, room)
		}

	}

	return m
}
