package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Name is for entities that have a visible name.
type Name string

func (Name) CTag() ecs.CID {
	return ecs.CID("Name")
}
