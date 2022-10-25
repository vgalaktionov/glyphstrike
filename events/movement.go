package events

import "github.com/vgalaktionov/roguelike-go/state"

type MovementAction struct {
	dX int
	dY int
}

func (ma *MovementAction) Handle(w *state.World) {
	w.Player.X += ma.dX
	w.Player.Y += ma.dY
}
