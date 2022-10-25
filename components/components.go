package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type Position struct {
	X, Y int
}

func (Position) ComponentTag() ecs.Tag {
	return ecs.Tag("Position")
}

type Renderable struct {
	Glyph rune
	Style tcell.Style
}

func (Renderable) ComponentTag() ecs.Tag {
	return ecs.Tag("Renderable")
}

type Player struct{}

func (Player) ComponentTag() ecs.Tag {
	return ecs.Tag("Player")
}

type LeftMover struct{}

func (LeftMover) ComponentTag() ecs.Tag {
	return ecs.Tag("LeftMover")
}
