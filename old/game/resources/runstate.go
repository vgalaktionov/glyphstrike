package resources

import "github.com/vgalaktionov/roguelike-go/ecs"

// RunState denotes the status of the game
type GameState int

const (
	PreRun GameState = iota
	PlayerTurn
	MonsterTurn
	AwaitingInput
)

func (GameState) RID() ecs.RID {
	return runStateID
}
