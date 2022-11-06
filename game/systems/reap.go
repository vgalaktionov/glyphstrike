package systems

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// Reap disposes of dead entities.
func Reap(w *ecs.World) {
	playerEnt := ecs.GetResource[resources.PlayerEntity](w).Entity()
	for _, e := range ecs.QueryEntitiesIter(w, CombatStats{}, Name("")) {
		stats := ecs.MustGetEntityComponent[CombatStats](w, e)

		if stats.HP <= 0 {
			if e == playerEnt {
				log.Print("You are dead.")
			} else {
				name := ecs.MustGetEntityComponent[Name](w, e)
				log.Printf("%s is dead.", name)
				ecs.RemoveEntity(w, e)
			}
		}
	}
}
