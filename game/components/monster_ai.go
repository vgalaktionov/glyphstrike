package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// MonsterAI marks a NPC monster with behavior.
type MonsterAI struct{}

func (MonsterAI) CTag() ecs.CID {
	return ecs.CID("MonsterAI")
}
