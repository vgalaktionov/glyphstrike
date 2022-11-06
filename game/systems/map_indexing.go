package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// MapIndexing deals with updating map state based on the ECS data
func MapIndexing(w *ecs.World) {
	gs := ecs.GetResource[resources.GameState](w)
	if gs == resources.AwaitingInput {
		return
	}
	m := ecs.GetResource[*resources.Map](w)
	m.PopulateBlocked()
	m.ClearContents()

	for _, e := range ecs.QueryEntitiesIter(w, Position{}, BlocksTile{}) {
		pos := ecs.MustGetEntityComponent[Position](w, e)

		if ecs.HasEntityComponent[BlocksTile](w, e) {
			m.BlockedTiles[pos.X][pos.Y] = true
		}

		m.TileContents[pos.X][pos.Y] = append(m.TileContents[pos.X][pos.Y], e)
	}
}
