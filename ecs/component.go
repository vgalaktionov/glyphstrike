package ecs

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

// GetEntityComponent retrieves a component by entity id.
// The result (if retrieved through the QueryEntities family of functions) is safe to cast to its intended type.
func GetEntityComponent[C Component](w *World, ent Entity) C {
	return w.components[(*new(C)).CTag()][ent].(C)
}

// SetEntityComponent replaces the given component for an entity.
func SetEntityComponent(w *World, c Component, ent Entity) {
	w.components[c.CTag()][ent] = c
}
