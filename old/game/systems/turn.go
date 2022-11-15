package systems

import (
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// UpdateTurn is the state machine for the next iteration of the game loop.
// The base loop alternates before player and monster turns, blocking on player input on their turn.
func UpdateTurn(w *ecs.World) {
	previousState := ecs.GetResource[resources.GameState](w)

	switch previousState {
	case resources.PreRun, resources.MonsterTurn:
		ecs.SetResource(w, resources.AwaitingInput)
	case resources.PlayerTurn:
		ecs.SetResource(w, resources.MonsterTurn)
	}
}
