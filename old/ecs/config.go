//go:build !js
// +build !js

package ecs

// avoid runtime reallocations for a reasonably sized game, can be tweaked later
const PreAllocateEntities = 100_000
const PreAllocateComponents = 100
