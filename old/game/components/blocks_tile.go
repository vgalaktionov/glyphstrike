package components

import "github.com/vgalaktionov/roguelike-go/ecs"

// BlocksTile marks an entity that cannot be passed through
type BlocksTile struct{}

func (BlocksTile) CID() ecs.CID {
	return blocksTileID
}
