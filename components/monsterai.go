package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type MonsterAI struct{}

const MonsterAITag = ecs.CTag("MonsterAI")

func (MonsterAI) CTag() ecs.CTag {
	return MonsterAITag
}
