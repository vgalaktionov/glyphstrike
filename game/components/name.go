package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Name is for entities that have a visible name.
type Name struct {
	Str string
}

func (Name) CTag() ecs.CTag {
	return ecs.CTag("Name")
}
