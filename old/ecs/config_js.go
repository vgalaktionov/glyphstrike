//go:build js
// +build js

package ecs

// WASM runtime does not appreciate our pre-allocation so keep it modest
const PreAllocateEntities = 1000
const PreAllocateComponents = 10
