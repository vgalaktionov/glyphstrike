package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// MapIndexing deals with updating map state based on the ECS data
func MapIndexing(w *ecs.World) {
	m := ecs.GetResource[resources.Map](w)
	m.PopulateBlocked()

	for e := range ecs.QueryEntitiesIter(w, Position{}, BlocksTile{}) {
		pos := ecs.GetEntityComponent[Position](w, e)
		m.BlockedTiles[pos.X][pos.Y] = true
	}
}
