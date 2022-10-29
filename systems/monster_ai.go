package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/events"
)

// ProcessMonsterAI simulates monster behaviour
func ProcessMonsterAI(w *ecs.World) {
	player := ecs.QueryEntitiesSingle(w, Player{}, Position{})
	playerPos := ecs.GetEntityComponent[Position](w, player)

	for e := range ecs.QueryEntitiesIter(w, Position{}, Viewshed{}, MonsterAI{}) {
		vs := ecs.GetEntityComponent[Viewshed](w, e)
		if vs.View.IsVisible(playerPos.X, playerPos.Y) {
			msg := "Monster shouts insults."
			ecs.DispatchEvent(w, events.ConsoleEvent{Message: msg})
		}
	}
}
