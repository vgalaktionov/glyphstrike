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
		e.Renderer.ShowCursor(0, 5)
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
	e.ECS.AddResource(m)
	// e.ECS.AddResource(systems.NewTestMap(mapX, mapY, mapX / 2, mapY / 2))

	playerX, playerY := m.Rooms[0].Center()
	e.ECS.AddEntity(
		components.Player{},
		components.Position{X: playerX, Y: playerY},
		components.Renderable{
			Glyph: '@',
			Style: tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorBlack),
		},
		components.Viewshed{Radius: 8, View: fov.New()},
	)

	e.Run()
}
