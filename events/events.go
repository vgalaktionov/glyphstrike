package events

import "github.com/vgalaktionov/roguelike-go/ecs"

type ConsoleEvent struct {
	Message string
}

const ConsoleEventTag = ecs.ETag("ConsoleEvent")

func (ConsoleEvent) ETag() ecs.ETag {
	return ConsoleEventTag
}
