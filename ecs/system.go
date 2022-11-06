package ecs

import "github.com/vgalaktionov/roguelike-go/util"

// System is just a function that acts on the ECS world.
type System func(*World)

// RegisterSystem adds a new blocking system to the ECS, to run on each tick of the game loop.
func RegisterSystem(w *World, s System) {
	w.systems = append(w.systems, s)
}

// RegisterSystem adds a new event system to the ECS, to run in a background goroutine.
func RegisterEventSystem(w *World, s System, events ...Event) {
	w.eventSystems = append(w.eventSystems, s)
	maxEvId := 0
	for _, ev := range events {
		maxEvId = util.MaxInt(int(ev.EID()), maxEvId)
	}
	if diff := maxEvId - len(w.resources) + 1; diff > 0 {
		w.events = append(w.events, make([]chan Event, diff)...)
	}
	for _, ev := range events {
		w.events[ev.EID()] = make(chan Event, 1000)
	}
}

// RunEventSystems is meant to be called once, before entering the main game loop.
func RunEventSystems(w *World) {
	for _, runEventSystem := range w.eventSystems {
		go runEventSystem(w)
	}
}

// RunSystems is meant to be called on every tick of the game loop.
func RunSystems(w *World) {
	for _, runSystem := range w.systems {
		runSystem(w)
	}
}
