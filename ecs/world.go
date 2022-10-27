package ecs

import (
	"github.com/vgalaktionov/roguelike-go/draw"
)

// Entity is an ID that can be used to retrieve associated components.
type Entity int

// CTag is the type of a component
type CTag string

// Component is a grouping of data used by one or multiple systems, that can be attached to an entity.
// Components are marked by a dummy function returning their type.
type Component interface {
	CTag() CTag
}

// RTag is the type of a resource
type RTag string

// Resource is a mutable singleton in the ECS, used for simplified special cases like the Map.
// Resources are marked by a dummy function returning their type.
type Resource interface {
	RTag() RTag
}

// ETag is the type of an event
type ETag string

// Event are messages, dispatched by any system, to communicate with background (event)systems.
// Events are marked by a dummy function returning their type.
type Event interface {
	ETag() ETag
}

// System is just a function that takes a renderer implementation and the ECS world.
type System func(draw.Renderer, *World)

// World encapsulates the internal datastructures of the ECS
type World struct {
	lastEntityID Entity
	entities     map[Entity]struct{}
	components   map[CTag]map[Entity]Component
	systems      []System
	resources    map[RTag]Resource
	events       map[ETag]chan Event
	eventSystems []System
}

// NewWorld returns an empty, usable world.
func NewWorld() *World {
	w := &World{
		0,
		make(map[Entity]struct{}),
		make(map[CTag]map[Entity]Component),
		nil,
		make(map[RTag]Resource),
		make(map[ETag]chan Event),
		nil,
	}
	return w
}

// AddEntity adds a new entity to the world with any number of attached components, and returns the ID.
func (w *World) AddEntity(components ...Component) Entity {
	w.lastEntityID++
	w.entities[w.lastEntityID] = struct{}{}
	for _, c := range components {
		c := c
		if w.components[c.CTag()] == nil {
			w.components[c.CTag()] = make(map[Entity]Component)
		}
		w.components[c.CTag()][w.lastEntityID] = c
	}
	return w.lastEntityID
}

// RemoveEntity removes an entity and all associated components by ID.
func (w *World) RemoveEntity(ent Entity) {
	delete(w.entities, ent)
	for _, components := range w.components {
		delete(components, ent)
	}
}

// AddResource registers a mutable singleton resource.
func (w *World) AddResource(r Resource) {
	w.resources[r.RTag()] = r
}

// GetResource retrieves a utable singleton resource by tag.
func (w *World) GetResource(tag RTag) Resource {
	return w.resources[tag]
}

// RegisterSystem adds a new blocking system to the ECS, to run on each tick of the game loop.
func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

// RegisterSystem adds a new event system to the ECS, to run in a background goroutine.
func (w *World) RegisterEventSystem(s System) {
	w.eventSystems = append(w.eventSystems, s)
}

// RegisterEvent sets up a buffered channel for the given event type.
func (w *World) RegisterEvent(e Event) {
	w.events[e.ETag()] = make(chan Event, 1000)
}

// GetEventChannel retrieves a channel by event tag.
func (w *World) GetEventChannel(tag ETag) chan Event {
	return w.events[tag]
}

// DispatchEvent sends an Event on the appropriate channel.
func (w *World) DispatchEvent(evt Event) {
	go func() {
		w.events[evt.ETag()] <- evt
	}()
}

// GetEntityComponent retries a component by tag and entity id.
// The result (if retrieved through the QueryEntities family of functions) is safe to cast to its intended type.
func (w *World) GetEntityComponent(tag CTag, ent Entity) Component {
	return w.components[tag][ent]
}

// SetEntityComponent replaces the given component for an entity.
func (w *World) SetEntityComponent(c Component, ent Entity) {
	w.components[c.CTag()][ent] = c
}

// EntityNotFound is a sentinel value for missing entities.
const EntityNotFound = Entity(-1)

// QueryEntitiesSingle takes templates (empty components) and returns the first entity with these components, or -1.
func (w *World) QueryEntitiesSingle(templates ...Component) Entity {

	for e := range w.entities {
		e := e
		hasAll := true
	inner:
		for _, t := range templates {
			t := t
			_, hasAll = w.components[t.CTag()][e]
			if !hasAll {
				break inner
			}
		}
		if hasAll {
			return e
		}
	}
	return EntityNotFound
}

// QueryEntitiesIter takes templates (empty components) and returns an iterable channel of entities with these components.
func (w *World) QueryEntitiesIter(templates ...Component) chan Entity {
	ch := make(chan Entity)
	go func() {
		defer close(ch)

		for e := range w.entities {
			hasAll := true
		inner:
			for _, t := range templates {
				_, hasAll = w.components[t.CTag()][e]
				if !hasAll {
					break inner
				}
			}
			if hasAll {
				ch <- e
			}
		}
	}()
	return ch
}

// RunEventSystems is meant to be called once, before entering the main game loop.
func (w *World) RunEventSystems(r draw.Renderer) {
	for _, runEventSystem := range w.eventSystems {
		go runEventSystem(r, w)
	}
}

// RunSystems is meant to be called on every tick of the game loop.
func (w *World) RunSystems(r draw.Renderer) {
	for _, runSystem := range w.systems {
		runSystem(r, w)
	}
}
