package systems

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/util"

	//lint:ignore ST1001 dot importing resources makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/game/resources"
)

// ClearMap only clears the map part of the screen, leaving UI elements intact.
func ClearMap(r draw.Screen) {
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
	r := ecs.GetResource[*Renderer](w)

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

// NewEmptyMap initializes a map with floor tiles.
func NewEmptyMap(mapWidth, mapHeight int) *Map {
	m := &Map{Tiles: make([][]TileType, mapWidth), Width: mapWidth, Height: mapHeight, RevealedTiles: make([][]bool, mapWidth)}
	for i := 0; i < mapWidth; i++ {
		m.Tiles[i] = make([]TileType, m.Height)
		m.RevealedTiles[i] = make([]bool, m.Height)
	}
	return m
}

// NewTestMap generates a map with random walls.
func NewTestMap(mapWidth, mapHeight int) *Map {
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
		m.Tiles[x][y] = WallTile

	}

	m.Rooms[0] = util.NewRect(0, 0, m.Width, m.Height)

	return m
}

// NewMapRoomsAndCorridors generates a map with rooms guaranteed to be connected by corridors.
func NewMapRoomsAndCorridors(mapWidth, mapHeight int) *Map {
	m := NewEmptyMap(mapWidth, mapHeight)
	m.Fill(WallTile)

	maxRooms := 30
	minSize := 6
	maxSize := 20

	for i := 0; i < maxRooms; i++ {
		// Roll for random room sizes within parameters.
		w := rand.Intn(maxSize-minSize) + minSize
		h := rand.Intn(maxSize-minSize) + minSize
		x := rand.Intn(m.Width - w - 1)
		y := rand.Intn(m.Height - h - 1)

		room := util.NewRect(x, y, w, h)
		ok := true
		// Don't overlap with other rooms
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
				// Flip a coin to decide which side our tunnels go
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
