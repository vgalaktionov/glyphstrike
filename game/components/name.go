package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Name is for entities that have a visible name.
type Name string

func (Name) CID() ecs.CID {
	return nameID
}
