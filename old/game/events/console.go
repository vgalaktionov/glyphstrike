package events

import "github.com/vgalaktionov/roguelike-go/ecs"

// ConsoleEvent represent an event dispatched when a system needs to log to the console.
type ConsoleEvent struct {
	Message string
}

func (ConsoleEvent) EID() ecs.EID {
	return consoleID
}
