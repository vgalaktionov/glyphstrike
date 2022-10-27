package resources

import (
	"math"

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
}

// InBounds returns whether an x,y coordinate pair is on the map.
func (m Map) InBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

// InBounds returns whether an x,y coordinate pair can be seen/moved through.
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
		y := y
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
	for x := int(math.Min(float64(x1), float64(x2))); x <= int(math.Max(float64(x1), float64(x2))); x++ {
		x := x
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

// ApplyVerticalTunnel draws a vertical tunnel between 2 points on an existing map.
func (m *Map) ApplyVerticalTunnel(y1, y2, x int) {
	for y := int(math.Min(float64(y1), float64(y2))); y <= int(math.Max(float64(y1), float64(y2))); y++ {
		y := y
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

const MapTag = ecs.RTag("Map")

func (Map) RTag() ecs.RTag {
	return MapTag
}
