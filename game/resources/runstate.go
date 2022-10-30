package resources

import "github.com/vgalaktionov/roguelike-go/ecs"

// RunState denotes the status of the game
type GameState int

const (
	PreRun GameState = iota
	PlayerTurn
	MonsterTurn
)

func (GameState) RTag() ecs.RTag {
	return ecs.RTag("GameState")
}
