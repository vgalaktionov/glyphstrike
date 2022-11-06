package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Player marks the player entity.
type Player struct{}

func (Player) CTag() ecs.CID {
	return ecs.CID("Player")
}
