package resources

import "github.com/vgalaktionov/roguelike-go/ecs"

type PlayerEntity ecs.Entity

func (p PlayerEntity) RID() ecs.RID {
	return playerID
}

func (p PlayerEntity) Entity() ecs.Entity {
	return ecs.Entity(p)
}
