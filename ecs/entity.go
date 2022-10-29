package ecs

// Entity is an ID that can be used to retrieve associated components.
type Entity int

// AddEntity adds a new entity to the world with any number of attached components, and returns the ID.
func AddEntity(w *World, components ...Component) Entity {
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
func RemoveEntity(w *World, ent Entity) {
	delete(w.entities, ent)
	for _, components := range w.components {
		delete(components, ent)
	}
}

// EntityNotFound is a sentinel value for missing entities.
const EntityNotFound = Entity(-1)

// QueryEntitiesSingle takes templates (empty components) and returns the first entity with these components, or -1.
func QueryEntitiesSingle(w *World, templates ...Component) Entity {

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
			return e
		}
	}
	return EntityNotFound
}

// QueryEntitiesIter takes templates (empty components) and returns an iterable channel of entities with these components.
func QueryEntitiesIter(w *World, templates ...Component) chan Entity {
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
