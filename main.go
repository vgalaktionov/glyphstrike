package main

import (
	"github.com/gdamore/tcell/v2"
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
	e.ECS.RegisterSystem(systems.LeftWalker)
	e.ECS.RegisterSystem(systems.Render)

	e.ECS.AddEntity(
		components.Player{},
		components.Position{X: 40, Y: 25},
		components.Renderable{
			Glyph: '@',
			Style: tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorBlack),
		},
	)

	for i := 0; i < 10; i++ {
		e.ECS.AddEntity(
			components.LeftMover{},
			components.Position{X: i * 7, Y: 20},
			components.Renderable{
				Glyph: '☺',
				Style: tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack),
			},
		)

	}

	e.Run()
}
