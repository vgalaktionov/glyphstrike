package ecs

import "fmt"

// CTag is the type of a component
type CTag string

// Component is a grouping of data used by one or multiple systems, that can be attached to an entity.
// Components are marked by a dummy function returning their type.
type Component interface {
	CTag() CTag
}

// HasEntityComponent checks for existence of a component by entity id.
func HasEntityComponent[C Component](w *World, ent Entity) bool {
	_, ok := w.components[(*new(C)).CTag()][ent]
	return ok
}

// MustGetEntityComponent retrieves a component by entity id.
// The result (if retrieved through the QueryEntities family of functions) is safe to cast to its intended type.
func MustGetEntityComponent[C Component](w *World, ent Entity) C {
	return w.components[(*new(C)).CTag()][ent].(C)
}

// GetEntityComponent retrieves a component by entity id, or returns an error if not found
func GetEntityComponent[C Component](w *World, ent Entity) (*C, error) {
	c, ok := w.components[(*new(C)).CTag()][ent].(C)
	if !ok {
		return nil, fmt.Errorf("no component %s found for entity %d", (*new(C)).CTag(), ent)
	}
	return &c, nil
}

// SetEntityComponent replaces the given component for an entity.
func SetEntityComponent(w *World, c Component, ent Entity) {
	if w.components[c.CTag()] == nil {
		w.components[c.CTag()] = make(map[Entity]Component)
	}
	w.components[c.CTag()][ent] = c
}

// RemoveEntityComponent deletes the given component for an entity.
func RemoveEntityComponent[C Component](w *World, ent Entity) {
	delete(w.components[(*new(C)).CTag()], ent)
}
