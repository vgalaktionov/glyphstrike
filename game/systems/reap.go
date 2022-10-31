package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
)

// Reap disposes of dead entities.
func Reap(w *ecs.World) {
	for e := range ecs.QueryEntitiesIter(w, CombatStats{}) {
		stats := ecs.MustGetEntityComponent[CombatStats](w, e)

		if stats.HP <= 0 {
			ecs.RemoveEntity(w, e)
		}
	}
}
