package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Position is an entity with a location on the map.
type Position struct {
	X, Y int
}

func (Position) CTag() ecs.CTag {
	return ecs.CTag("Position")
}
