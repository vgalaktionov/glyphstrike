package systems

import (
	"log"

	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
	"github.com/vgalaktionov/roguelike-go/util"
)

// ProcessMonsterAI simulates monster behaviour
func ProcessMonsterAI(w *ecs.World) {
	gs := ecs.GetResource[resources.GameState](w)
	if gs != resources.MonsterTurn {
		return
	}

	player, err := ecs.QueryEntitiesSingle(w, Player{}, Position{})
	// If we don't have a player entity, we should bail out (until death is implemented)
	if err != nil {
		log.Fatalln("no player found")
	}
	playerPos := ecs.GetEntityComponent[Position](w, player)
	m := ecs.GetResource[*resources.Map](w)

	for e := range ecs.QueryEntitiesIter(w, Position{}, Viewshed{}, MonsterAI{}, Name("")) {
		vs := ecs.GetEntityComponent[Viewshed](w, e)
		name := ecs.GetEntityComponent[Name](w, e)
		pos := ecs.GetEntityComponent[Position](w, e)

		dist := util.Distance(pos.X, pos.Y, playerPos.X, playerPos.Y)
		if dist < 1.5 {
			// attack goes here
			log.Printf("%s shouts insults.", name)
			continue
		}

		if vs.View.IsVisible(playerPos.X, playerPos.Y) {
			g := m.GetGridFor(pos.X, pos.Y)
			path := g.GetPath(float64(pos.X), float64(pos.Y), float64(playerPos.X), float64(playerPos.Y), true, true)
			if path == nil || path.Length() < 2 {
				continue
			}
			if nextCell := path.Next(); nextCell != nil && !(nextCell.X == playerPos.X && nextCell.Y == playerPos.Y) {
				ecs.SetEntityComponent(w, Position{X: nextCell.X, Y: nextCell.Y}, e)
			}
		}
	}
}
