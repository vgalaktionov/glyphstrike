package systems

import (
	"os"

	"github.com/gdamore/tcell/v2"

	//lint:ignore ST1001 sheesh
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

type LeftWalker struct{}

func (LeftWalker) Query() []ecs.Tag {
	return []ecs.Tag{Position{}.ComponentTag(), LeftMover{}.ComponentTag()}
}

func (LeftWalker) Process(s tcell.Screen, components ...*ecs.Component) {
	pos := (*components[0]).(Position)
	pos.X--
	maxX, _ := s.Size()
	if pos.X > maxX {
		pos.X = 0
	}
}

type PlayerInput struct{}

func (PlayerInput) Query() []ecs.Tag {
	return []ecs.Tag{Player{}.ComponentTag(), Position{}.ComponentTag()}
}

func (PlayerInput) Process(s tcell.Screen, components ...*ecs.Component) {
	position := (*components[1]).(*Position)
	event := s.PollEvent()
	switch ev := event.(type) {
	case *tcell.EventResize:
		s.Sync()

	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			s.Clear()
			s.ShowCursor(0, 0)
			s.Fini()
			os.Exit(0)
		}

		switch ev.Key() {
		case tcell.KeyLeft:
			position.X--
		case tcell.KeyRight:
			position.X++
		case tcell.KeyUp:
			position.Y--
		case tcell.KeyDown:
			position.Y++
		}

	case *tcell.EventMouse:
		// mouseX, mouseY := ev.Position()

		// 		switch ev.Buttons() {
		// 		case tcell.Button1, tcell.Button2:
		// 		case tcell.ButtonNone:
		// 		}
		// 	}
		// }
	}
}

type Render struct{}

func (Render) Query() []ecs.Tag {
	return []ecs.Tag{Position{}.ComponentTag(), Renderable{}.ComponentTag()}
}

func (Render) Process(s tcell.Screen, components ...*ecs.Component) {
	pos := (*components[0]).(*Position)
	renderable := (*components[1]).(*Renderable)

	s.SetContent(pos.X, pos.Y, renderable.Glyph, nil, draw.DEFAULT_STYLE)
}
