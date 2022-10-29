package ecs

// RTag is the type of a resource
type RTag string

// Resource is a mutable singleton in the ECS, used for simplified special cases like the Map.
// Resources are marked by a dummy function returning their type.
type Resource interface {
	RTag() RTag
}


// AddResource registers a mutable singleton resource.
func AddResource(w *World, r Resource) {
	w.resources[r.RTag()] = r
}

// GetResource retrieves a singleton resource by tag.
func GetResource[R Resource](w *World) R {
	return w.resources[(*new(R)).RTag()].(R)
}