package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/norendren/go-fov/fov"
	"github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/engine"
	"github.com/vgalaktionov/roguelike-go/systems"
)

func main() {
	e := engine.NewEngine()

	quit := func() {
		maybePanic := recover()
		e.Renderer.Clear()
		e.Renderer.ShowCursor(0, 0)
		e.Renderer.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	e.ECS.RegisterSystem(systems.PlayerInput)
	e.ECS.RegisterSystem(systems.Visibility)
	e.ECS.RegisterSystem(systems.RenderMap)
	e.ECS.RegisterSystem(systems.Render)

	mapX, mapY := e.Renderer.Size()

	m := systems.NewMapRoomsAndCorridors(mapX, mapY)
	// m := systems.NewTestMap(mapX, mapY)
	e.ECS.AddResource(m)

	playerX, playerY := m.Rooms[0].Center()
	e.ECS.AddEntity(
		components.Player{},
		components.Position{X: playerX, Y: playerY},
		components.Renderable{
			Glyph: '⊛',
			Style: tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack.TrueColor()),
		},
		components.Viewshed{Radius: 8, View: fov.New()},
	)

	for i := 1; i < len(m.Rooms); i++ {
		x, y := m.Rooms[i].Center()
		e.ECS.AddEntity(
			components.Position{X: x, Y: y},
			components.Renderable{Glyph: 'g', Style: tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor())},
			components.Viewshed{Radius: 8, View: fov.New()},
		)
	}

	e.Run()
}
