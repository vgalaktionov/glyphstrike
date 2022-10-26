package resources

import "github.com/vgalaktionov/roguelike-go/ecs"

type TileType int

const (
	FloorTile TileType = iota
	WallTile
)

type Map [][]TileType

const MapTag = ecs.RTag("Map")

func (Map) RTag() ecs.RTag {
	return MapTag
}
