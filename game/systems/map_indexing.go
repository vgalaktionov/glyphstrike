package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// MapIndexing deals with updating map state based on the ECS data
func MapIndexing(w *ecs.World) {
	m := ecs.GetResource[*resources.Map](w)
	m.PopulateBlocked()
	m.ClearContents()

	for e := range ecs.QueryEntitiesIter(w, Position{}, BlocksTile{}) {
		pos := ecs.MustGetEntityComponent[Position](w, e)

		if ecs.HasEntityComponent[BlocksTile](w, e) {
			m.BlockedTiles[pos.X][pos.Y] = true
		}

		m.TileContents[pos.X][pos.Y] = append(m.TileContents[pos.X][pos.Y], e)
	}
}
