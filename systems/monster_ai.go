package systems

import (

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/events"
)

func ProcessMonsterAI(r draw.Renderer, w *ecs.World) {
	cec := w.GetEventChannel(events.ConsoleEventTag)
	for range w.QueryEntitiesIter(Position{}, Viewshed{}, MonsterAI{}) {
		msg := "Monsters consider their own existence"
		go func() {
			cec <- events.ConsoleEvent{Message: msg}
		}()
	}
}
