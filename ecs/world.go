package ecs

// World encapsulates the internal datastructures of the ECS.
// World does not have methods, instead it is operated on by functions from the `ecs` package.
//
// This is a stylistic tradeoff that presents a nice API on the caller side without requiring typecasting
// and interfacing with the _Tag strings directly. However due to limitations in the go generics implementation,
// we are not allowed to introduce generic parameters in methods without making the struct and ALL references to
// the type require type params. Skipping methods and using functions instead allows us to sidestep this issue.
type World struct {
	lastEntityID Entity
	entities     map[Entity]struct{}
	components   map[CTag]map[Entity]Component
	systems      []System
	resources    map[RTag]Resource
	events       map[ETag]chan Event
	eventSystems []System
}

// NewWorld returns an empty, usable world.
func NewWorld() *World {
	w := &World{
		0,
		make(map[Entity]struct{}),
		make(map[CTag]map[Entity]Component),
		nil,
		make(map[RTag]Resource),
		make(map[ETag]chan Event),
		nil,
	}
	return w
}
