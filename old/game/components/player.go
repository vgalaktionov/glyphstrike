package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Player marks the player entity.
type Player struct{}

func (Player) CID() ecs.CID {
	return playerID
}
