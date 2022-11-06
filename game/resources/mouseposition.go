package resources

import "github.com/vgalaktionov/roguelike-go/ecs"

type MousePosition struct {
	X, Y  int
	Moved bool
}

func (MousePosition) RID() ecs.RID {
	return mousePositionID
}
