package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// WantsToMelee marks the intent to engage another entity in melee combat.
type WantsToMelee struct {
	Target ecs.Entity
}

func (WantsToMelee) CID() ecs.CID {
	return wantsToMeleeID
}
