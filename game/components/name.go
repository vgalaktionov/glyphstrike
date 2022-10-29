package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Name is for entities that have a visible name.
type Name struct {
	Str string
}

const NameTag = ecs.CTag("Name")

func (Name) CTag() ecs.CTag {
	return NameTag
}
