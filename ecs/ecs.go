package ecs

import (
	set "github.com/deckarep/golang-set/v2"
	"github.com/gdamore/tcell/v2"
)

type World struct {
	lastEntityID Entity
	entities     set.Set[Entity]
	components   map[Tag]map[Entity]*Component
	systems      []System
}

func (w *World) AddEntity(components ...Component) Entity {
	w.lastEntityID++
	w.entities.Add(w.lastEntityID)
	for _, c := range components {
		if w.components[c.ComponentTag()] == nil {
			w.components[c.ComponentTag()] = make(map[Entity]*Component)
		}
		w.components[c.ComponentTag()][w.lastEntityID] = &c
	}
	return w.lastEntityID
}

func (w *World) RemoveEntity(e Entity) {
	w.entities.Remove(e)
	for _, components := range w.components {
		delete(components, e)
	}
}

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) RunQuery(tags []Tag) [][]*Component {
	entities := w.entities
	for _, tag := range tags {
		tagEntities := set.NewSet[Entity]()
		for e := range w.components[tag] {
			tagEntities.Add(e)
		}
		entities = entities.Intersect(tagEntities)
	}

	components := make([][]*Component, entities.Cardinality())
	for i, e := range entities.ToSlice() {
		for _, tag := range tags {
			components[i] = append(components[i], w.components[tag][e])
		}
	}
	return components
}

func (w *World) Process(screen tcell.Screen) {
	for _, sys := range w.systems {

		for _, components := range w.RunQuery(sys.Query()) {
			sys.Process(screen, components...)
		}
	}
}

type Entity int

type Tag string

type Component interface {
	ComponentTag() Tag
}

type System interface {
	Query() []Tag
	Process(tcell.Screen, ...*Component)
}

var ScreenWidth, ScreenHeight int

func NewWorld() *World {
	w := &World{0, set.NewSet[Entity](), make(map[Tag]map[Entity]*Component), nil}
	return w
}
