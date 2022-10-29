package events

import "github.com/vgalaktionov/roguelike-go/ecs"

// ConsoleEvent represent an event dispatched when a system needs to log to the console.
type ConsoleEvent struct {
	Message string
}

const ConsoleEventTag = ecs.ETag("ConsoleEvent")

func (ConsoleEvent) ETag() ecs.ETag {
	return ConsoleEventTag
}
