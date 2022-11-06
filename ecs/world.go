package ecs

import (
	"github.com/bits-and-blooms/bitset"
)

// World encapsulates the internal datastructures of the ECS.
// World does not have methods, instead it is operated on by functions from the `ecs` package.
//
// This is a stylistic tradeoff that presents a nice API on the caller side without requiring typecasting
// and interfacing with the _Tag strings directly. However due to limitations in the go generics implementation,
// we are not allowed to introduce generic parameters in methods without making the struct and ALL references to
// the type require type params. Skipping methods and using functions instead allows us to sidestep this issue.
type World struct {
	lastEntityID     Entity
	entities         []Entity
	componentIndices []*bitset.BitSet
	components       [][]Component
	systems          []System
	resources        []Resource
	events           []chan Event
	eventSystems     []System
}

// avoid runtime reallocations for a reasonably sized game, can be tweaked later
const PreAllocateEntities = 100_000
const PreAllocateComponents = 100

// NewWorld returns an empty, usable world.
func NewWorld() *World {
	w := &World{
		0,
		make([]Entity, 0, PreAllocateEntities),
		make([]*bitset.BitSet, PreAllocateComponents),
		make([][]Component, PreAllocateComponents),
		nil,
		[]Resource{},
		[]chan Event{},
		nil,
	}
	for i := 0; i < len(w.componentIndices); i++ {
		w.componentIndices[i] = bitset.New(PreAllocateEntities)
	}
	for i := 0; i < len(w.components); i++ {
		w.components[i] = make([]Component, PreAllocateEntities)
	}
	return w
}
