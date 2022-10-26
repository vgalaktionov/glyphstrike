package resources

import (
	"math"

	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type TileType int

const (
	FloorTile TileType = iota
	WallTile
)

type Map struct {
	Tiles  [][]TileType
	Width  int
	Height int
}

func (m *Map) InBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

func (m *Map) IsOpaque(x, y int) bool {
	return m.Tiles[x][y] == WallTile
}

func (m *Map) Fill(tt TileType) {
	for x, col := range m.Tiles {
		for y := range col {
			m.Tiles[x][y] = tt
		}
	}
}

func (m *Map) ApplyRoom(room draw.Rect) {
	for y := room.Y1 + 1; y <= room.Y2; y++ {
		for x := room.X1 + 1; x <= room.X2; x++ {
			m.Tiles[x][y] = FloorTile
		}
	}
}

func (m *Map) ApplyHorizontalTunnel(x1, x2, y int) {
	for x := int(math.Min(float64(x1), float64(x2))); y <= int(math.Max(float64(x1), float64(x2))); x++ {
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

func (m *Map) ApplyVerticalTunnel(y1, y2, x int) {
	for y := int(math.Min(float64(y1), float64(y2))); y <= int(math.Max(float64(y1), float64(y2))); y++ {
		if m.InBounds(x, y) {
			m.Tiles[x][y] = FloorTile
		}
	}
}

const MapTag = ecs.RTag("Map")

func (Map) RTag() ecs.RTag {
	return MapTag
}
