package ecs

import (
	set "github.com/deckarep/golang-set/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
)

type Entity int

type CTag string

type Component interface {
	CTag() CTag
}

type RTag string

type Resource interface {
	RTag() RTag
}

type ETag string

type Event interface {
	ETag() ETag
}

type System func(draw.Renderer, *World)

type World struct {
	lastEntityID Entity
	entities     set.Set[Entity]
	components   map[CTag]map[Entity]Component
	systems      []System
	resources    map[RTag]Resource
	events       map[ETag]chan Event
	eventSystems []System
}

func NewWorld() *World {
	w := &World{
		0,
		set.NewSet[Entity](),
		make(map[CTag]map[Entity]Component),
		nil,
		make(map[RTag]Resource),
		make(map[ETag]chan Event),
		nil,
	}
	return w
}

func (w *World) AddEntity(components ...Component) Entity {
	w.lastEntityID++
	w.entities.Add(w.lastEntityID)
	for _, c := range components {
		c := c
		if w.components[c.CTag()] == nil {
			w.components[c.CTag()] = make(map[Entity]Component)
		}
		w.components[c.CTag()][w.lastEntityID] = c
	}
	return w.lastEntityID
}

func (w *World) RemoveEntity(ent Entity) {
	w.entities.Remove(ent)
	for _, components := range w.components {
		delete(components, ent)
	}
}

func (w *World) AddResource(r Resource) {
	w.resources[r.RTag()] = r
}

func (w *World) GetResource(tag RTag) Resource {
	return w.resources[tag]
}

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) RegisterEventSystem(s System) {
	w.eventSystems = append(w.eventSystems, s)
}

func (w *World) RegisterEvent(e Event) {
	w.events[e.ETag()] = make(chan Event, 1000)
}

func (w *World) GetEventChannel(tag ETag) chan Event {
	return w.events[tag]
}

func (w *World) DispatchEvent(evt Event) {
	go func() {
		w.events[evt.ETag()] <- evt
	}()
}

func (w *World) GetEntityComponent(tag CTag, ent Entity) Component {
	return w.components[tag][ent]
}

// SetEntityComponent replaces the given component for an entity.
func (w *World) SetEntityComponent(c Component, ent Entity) {
	w.components[c.CTag()][ent] = c
}

const EntityNotFound = Entity(-1)

// QueryEntitiesSingle takes templates (empty components) and returns the first entity with these components, or -1.
func (w *World) QueryEntitiesSingle(templates ...Component) Entity {

	for e := range w.entities.Iter() {
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

		for e := range w.entities.Iter() {
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

func (w *World) RunEventSystems(r draw.Renderer) {
	for _, runEventSystem := range w.eventSystems {
		go runEventSystem(r, w)
	}
}

func (w *World) RunSystems(r draw.Renderer) {
	for _, runSystem := range w.systems {
		runSystem(r, w)
	}
}
