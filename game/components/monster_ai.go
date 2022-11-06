package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// MonsterAI marks a NPC monster with behavior.
type MonsterAI struct{}

func (MonsterAI) CID() ecs.CID {
	return monsterAiID
}
