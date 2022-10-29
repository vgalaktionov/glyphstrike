package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Position is an entity with a location on the map.
type Position struct {
	X, Y int
}

const PositionTag = ecs.CTag("Position")

func (Position) CTag() ecs.CTag {
	return PositionTag
}
