package systems

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/util"
)

// MeleeCombat handles all melee combat interactions between entities and queues resolved damage to be applied.
func MeleeCombat(w *ecs.World) {
	for _, e := range ecs.QueryEntitiesIter(w, WantsToMelee{}, Name(""), CombatStats{}) {
		stats := ecs.MustGetEntityComponent[CombatStats](w, e)
		target := ecs.MustGetEntityComponent[WantsToMelee](w, e)
		if stats.HP > 0 {
			targetStats := ecs.MustGetEntityComponent[CombatStats](w, target.Target)
			if targetStats.HP > 0 {
				name := ecs.MustGetEntityComponent[Name](w, e)
				targetName := ecs.MustGetEntityComponent[Name](w, target.Target)
				damage := util.MaxInt(0, stats.Power-targetStats.Defense)

				if damage <= 0 {
					log.Printf("%s is unable to hurt %s.", name, targetName)
				} else {
					log.Printf("%s hits %s, for %d hp.", name, targetName, damage)
					NewDamage(w, damage, target.Target)
				}
			}
		}
		ecs.RemoveEntityComponent[WantsToMelee](w, e)
	}
}
