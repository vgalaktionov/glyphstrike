package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// MonsterAI marks a NPC monster with behavior.
type MonsterAI struct{}

const MonsterAITag = ecs.CTag("MonsterAI")

func (MonsterAI) CTag() ecs.CTag {
	return MonsterAITag
}
