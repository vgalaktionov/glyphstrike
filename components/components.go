package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Position struct {
	X, Y int
}

const PositionTag = ecs.CTag("Position")

func (Position) CTag() ecs.CTag {
	return PositionTag
}

type Renderable struct {
	Glyph rune
	Style tcell.Style
}

const RenderableTag = ecs.CTag("Renderable")

func (Renderable) CTag() ecs.CTag {
	return RenderableTag
}

type Player struct{}

const PlayerTag = ecs.CTag("Player")

func (Player) CTag() ecs.CTag {
	return PlayerTag
}

type LeftMover struct{}

const LeftMoverTag = ecs.CTag("LeftMover")

func (LeftMover) CTag() ecs.CTag {
	return LeftMoverTag
}
