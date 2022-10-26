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

type System func(draw.Renderer, *World)

type World struct {
	lastEntityID Entity
	entities     set.Set[Entity]
	components   map[CTag]map[Entity]*Component
	systems      []System
}

func NewWorld() *World {
	w := &World{0, set.NewSet[Entity](), make(map[CTag]map[Entity]*Component), nil}
	return w
}

func (w *World) AddEntity(components ...Component) Entity {
	w.lastEntityID++
	w.entities.Add(w.lastEntityID)
	for _, c := range components {
		c := c
		if w.components[c.CTag()] == nil {
			w.components[c.CTag()] = make(map[Entity]*Component)
		}
		w.components[c.CTag()][w.lastEntityID] = &c
	}
	return w.lastEntityID
}

func (w *World) RemoveEntity(ent Entity) {
	w.entities.Remove(ent)
	for _, components := range w.components {
		delete(components, ent)
	}
}

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) GetEntityComponent(tag CTag, ent Entity) Component {
	return *w.components[tag][ent]
}

// SetEntityComponent replaces the given component for an entity.
func (w *World) SetEntityComponent(c Component, ent Entity) {
	w.components[c.CTag()][ent] = &c
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

func (w *World) RunSystems(r draw.Renderer) {
	for _, runSystem := range w.systems {
		runSystem(r, w)
	}
}
