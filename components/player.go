package components

import "github.com/vgalaktionov/roguelike-go/ecs"

type Player struct{}

const PlayerTag = ecs.CTag("Player")

func (Player) CTag() ecs.CTag {
	return PlayerTag
}
