package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// Player marks the player entity.
type Player struct{}

const PlayerTag = ecs.CTag("Player")

func (Player) CTag() ecs.CTag {
	return PlayerTag
}
