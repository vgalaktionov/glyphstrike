package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// CombatStats are possessed by everything that can engage in combat.
type CombatStats struct {
	MaxHP   int
	HP      int
	Defense int
	Power   int
}

func (CombatStats) CID() ecs.CID {
	return combatStatsID
}
