package ecs

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

// CID is the type of a component
type CID uint32

// Component is a grouping of data used by one or multiple systems, that can be attached to an entity.
// Components are marked by a dummy function returning their type.
type Component interface {
	CID() CID
}

// HasEntityComponent checks for existence of a component by entity id.
func HasEntityComponent[C Component](w *World, e Entity) bool {
	cid := (*new(C)).CID()
	if int(cid) > len(w.componentIndices) {
		return false
	}
	return w.componentIndices[cid].Test(uint(e))
}

// MustGetEntityComponent retrieves a component by entity id, or nil
// The result (if retrieved through the QueryEntities family of functions) is safe to cast to its intended type.
func MustGetEntityComponent[C Component](w *World, e Entity) C {
	return w.components[(*new(C)).CID()][e].(C)
}

// GetEntityComponent retrieves a component by entity id, or returns an error if not found
func GetEntityComponent[C Component](w *World, e Entity) (C, error) {
	cid := (*new(C)).CID()
	var nilComp C
	if int(cid)+1 > len(w.componentIndices) {
		return nilComp, fmt.Errorf("no component %d found", cid)
	}
	if !w.componentIndices[cid].Test(uint(e)) {
		return nilComp, fmt.Errorf("no component %d found for entity %d", cid, e)
	}
	return w.components[cid][e].(C), nil
}

// SetEntityComponent replaces the given component for an entity.
func SetEntityComponent(w *World, c Component, e Entity) {
	if w.components[c.CID()] == nil {
		w.components[c.CID()] = make([]Component, e+1)
	}
	if w.componentIndices[c.CID()] == nil {
		w.componentIndices[c.CID()] = bitset.New(uint(e))
	}
	if diff := int(e) - len(w.components[c.CID()]) + 1; diff > 0 {
		w.components[c.CID()] = append(w.components[c.CID()], make([]Component, diff)...)
	}
	w.components[c.CID()][e] = c
	w.componentIndices[c.CID()].Set(uint(e))
}

// RemoveEntityComponent deletes the given component for an entity.
func RemoveEntityComponent[C Component](w *World, e Entity) {
	cid := (*new(C)).CID()
	if len(w.components)+1 > int(e) {
		w.components[cid][e] = nil
	}
	if len(w.componentIndices)+1 > int(e) {
		w.componentIndices[cid] = w.componentIndices[cid].Clear(uint(e))
	}
}
