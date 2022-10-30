package resources

import (
	"math/rand"

	"github.com/solarlune/paths"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/util"
)

// TileType represents a tile with associated rendering and behavior properties
type TileType int

const (
	FloorTile TileType = iota
	WallTile
)

// Map represents a renderable stateful game map.
type Map struct {
	Tiles         [][]TileType
	Width         int
	Height        int
	Rooms         []util.Rect
	RevealedTiles [][]bool
	BlockedTiles  [][]bool
}

// NewEmptyMap initializes a map with floor tiles.
func NewEmptyMap(mapWidth, mapHeight int) *Map {
	m := &Map{
		Tiles: make([][]TileType, mapWidth),
		Width: mapWidth, Height: mapHeight,
		RevealedTiles: make([][]bool, mapWidth),
		BlockedTiles:  make([][]bool, mapWidth),
	}
	for i := 0; i < mapWidth; i++ {
		m.Tiles[i] = make([]TileType, m.Height)
		m.RevealedTiles[i] = make([]bool, m.Height)
		m.BlockedTiles[i] = make([]bool, m.Height)
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

func (m *Map) PopulateBlocked() {
	for x := range m.BlockedTiles {
		for y := range m.BlockedTiles[x] {
			m.BlockedTiles[x][y] = m.Tiles[x][y] == WallTile
		}
	}
}

// InBounds returns whether an x,y coordinate pair is on the map.
func (m Map) InBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

// IsOpaque returns whether an x,y coordinate pair can be seen/moved through.
func (m Map) IsOpaque(x, y int) bool {
	return m.Tiles[x][y] == WallTile
}

// Fill fills the map with a single tile.
func (m *Map) Fill(tt TileType) {
	for x, col := range m.Tiles {
		for y := range col {
			m.Tiles[x][y] = tt
		}
	}
}

// ApplyRoom draws a rectangular room on an existing map.
func (m *Map) ApplyRoom(room util.Rect) {
	for y := room.Y1 + 1; y <= room.Y2; y++ {
		for x := room.X1 + 1; x <= room.X2; x++ {
			x := x
			if m.InBounds(x, y) {
				m.Tiles[x][y] = FloorTile
			}
		}
	}
}

// ApplyHorizontalTunnel draws a horizontal tunnel between 2 points on an existing map.
func (m *Map) ApplyHorizontalTunnel(x1, x2, y int) {
	for x := util.MinInt(x1, x2); x <= util.MaxInt(x1, x2); x++ {
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

// ApplyVerticalTunnel draws a vertical tunnel between 2 points on an existing map.
func (m *Map) ApplyVerticalTunnel(y1, y2, x int) {
	for y := util.MinInt(y1, y2); y <= util.MaxInt(y1, y2); y++ {
		y := y
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

func (m *Map) GetGrid() *paths.Grid {
	g := paths.NewGrid(m.Width, m.Height, 1, 1)

	for _, c := range g.AllCells() {
		c.Walkable = !m.IsOpaque(c.X, c.Y)
	}

	return g
}

const MapTag = ecs.RTag("Map")

func (*Map) RTag() ecs.RTag {
	return MapTag
}
