package ecs

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/bits-and-blooms/bitset"
	"github.com/vgalaktionov/roguelike-go/util"
)

// Entity is an ID that can be used to retrieve associated components.
type Entity uint

// An entity has not been found. (if we run out of uints we have bigger challenges than this bug)
const MissingEntity = Entity(math.MaxUint)

// AddEntity adds a new entity to the world with any number of attached components, and returns the ID.
func AddEntity(w *World, components ...Component) Entity {
	eid := w.lastEntityID
	w.entities = append(w.entities, eid)
	maxCid := 0
	for _, c := range components {
		maxCid = util.MaxInt(int(c.CID()), maxCid)
	}
	if diff := maxCid - len(w.components) + 1; diff >= 0 {
		w.components = append(w.components, make([][]Component, diff)...)
		w.componentIndices = append(w.componentIndices, make([]*bitset.BitSet, diff)...)
	}
	for _, c := range components {
		SetEntityComponent(w, c, eid)
	}
	w.lastEntityID++
	return eid
}

// HasEntity checks if a given entity exists (mostly useful for testing)
func HasEntity(w *World, e Entity) bool {
	if len(w.entities) <= int(e) {
		return false
	}
	return w.entities[e] != MissingEntity
}

// RemoveEntity removes an entity and all associated components by ID.
func RemoveEntity(w *World, e Entity) {
	if len(w.entities) >= int(e)+1 {
		w.entities[e] = MissingEntity
	}
	for cid := range w.components {
		if len(w.components[cid]) >= int(e)+1 {
			w.components[cid][e] = nil
		}
		w.componentIndices[cid].Clear(uint(e))
	}
}

// QueryEntitiesSingle takes templates (empty components) and returns the first entity with these components, or error.
func QueryEntitiesSingle(w *World, templates ...Component) (Entity, error) {
	if len(templates) == 0 {
		if len(w.entities) > 0 {
			return w.entities[0], nil
		} else {
			return MissingEntity, fmt.Errorf("no entities have been added to the world")
		}
	}
	filter := w.componentIndices[templates[0].CID()].Clone()
	for _, t := range templates[1:] {
		tcid := t.CID()
		if int(tcid)+1 > len(w.componentIndices) {
			return MissingEntity, fmt.Errorf("no component %d has been added to the world", tcid)
		}
		filter.InPlaceIntersection(w.componentIndices[tcid])
	}
	next, ok := filter.NextSet(0)
	if !ok {
		return MissingEntity, fmt.Errorf("no entity found with components +%v", templates)
	}
	return Entity(next), nil

}

// QueryEntitiesIter takes templates (empty components) and returns a slice of entities with these components.
func QueryEntitiesIter(w *World, templates ...Component) []Entity {
	if len(templates) == 0 {
		return w.entities
	}

	filter := w.componentIndices[templates[0].CID()].Clone()
	for _, t := range templates[1:] {
		filter.InPlaceIntersection(w.componentIndices[t.CID()])
	}
	out := make([]uint, filter.Count())
	filter.NextSetMany(0, out)
	// it's already a slice of uints, let's not unnecessarily loop
	return *(*[]Entity)(unsafe.Pointer(&out))
}
