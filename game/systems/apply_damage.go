package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
)

// ApplyDamage manages the application of damage from all sources to entities.
func ApplyDamage(w *ecs.World) {
	for e := range ecs.QueryEntitiesIter(w, SufferDamage{}, CombatStats{}) {
		sd := ecs.MustGetEntityComponent[SufferDamage](w, e)
		stats := ecs.MustGetEntityComponent[CombatStats](w, e)

		for _, dmg := range sd.Damage {
			stats.HP -= dmg
		}
		ecs.SetEntityComponent(w, SufferDamage{}, e)
		ecs.SetEntityComponent(w, stats, e)
	}
}
