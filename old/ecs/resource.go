package ecs

// RID is the type of a resource
type RID uint

// Resource is a mutable singleton in the ECS, used for simplified special cases like the Map.
// Resources are marked by a dummy function returning their type.
type Resource interface {
	RID() RID
}

// SetResource adds a "global" singleton resource.
func AddResource(w *World, r Resource) {
	if diff := int(r.RID()) - len(w.resources) + 1; diff > 0 {
		w.resources = append(w.resources, make([]Resource, diff)...)
	}
	SetResource(w, r)
}

// SetResource sets a "global" singleton resource.
func SetResource(w *World, r Resource) {
	w.resources[r.RID()] = r
}

// GetResource retrieves a singleton resource by tag.
func GetResource[R Resource](w *World) R {
	return w.resources[(*new(R)).RID()].(R)
}
