package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/norendren/go-fov/fov"
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

type Viewshed struct {
	Radius int
	View   *fov.View
}

const ViewshedTag = ecs.CTag("Viewshed")

func (Viewshed) CTag() ecs.CTag {
	return ViewshedTag
}

type MonsterAI struct{}

const MonsterAITag = ecs.CTag("MonsterAI")

func (MonsterAI) CTag() ecs.CTag {
	return MonsterAITag
}
