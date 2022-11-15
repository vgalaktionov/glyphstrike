package components

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
)

// SufferDamage tracks the damage to be applied to a combat entity.
type SufferDamage struct {
	Damage []int
}

func (SufferDamage) CID() ecs.CID {
	return sufferDamageID
}

func NewDamage(w *ecs.World, dmg int, e ecs.Entity) {
	var existingDamage SufferDamage
	existingDamage, err := ecs.GetEntityComponent[SufferDamage](w, e)
	if err != nil {
		existingDamage = SufferDamage{}
	}
	newDamage := SufferDamage{Damage: append(existingDamage.Damage, dmg)}
	ecs.SetEntityComponent(w, newDamage, e)
}
