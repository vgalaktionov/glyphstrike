package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// BlocksTile marks an entity that cannot be passed through
type BlocksTile struct{}

func (BlocksTile) CTag() ecs.CID {
	return ecs.CID("BlocksTile")
}
