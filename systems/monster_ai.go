package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/events"
)

// ProcessMonsterAI simulates monster behaviour
func ProcessMonsterAI(w *ecs.World) {
	player := w.QueryEntitiesSingle(Player{}, Position{})
	playerPos := w.GetEntityComponent(PositionTag, player).(Position)

	for e := range w.QueryEntitiesIter(Position{}, Viewshed{}, MonsterAI{}) {
		vs := w.GetEntityComponent(ViewshedTag, e).(Viewshed)
		if vs.View.IsVisible(playerPos.X, playerPos.Y) {
			msg := "Monster shouts insults."
			w.DispatchEvent(events.ConsoleEvent{Message: msg})
		}
	}
}
