package systems

import (
	"os"

	"github.com/gdamore/tcell/v2"

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

func LeftWalker(r draw.Renderer, w *ecs.World) {
	for e := range w.QueryEntitiesIter(LeftMover{}, Position{}) {
		pos := w.GetEntityComponent(PositionTag, e).(Position)
		pos.X--
		maxX, _ := r.Size()
		if pos.X < 0 {
			pos.X = maxX
		}
		w.SetEntityComponent(pos, e)
	}
}

func PlayerInput(r draw.Renderer, w *ecs.World) {
	event := r.PollEvent()
	for e := range w.QueryEntitiesIter(Player{}, Position{}) {
		pos := w.GetEntityComponent(Position{}.CTag(), e).(Position)

		switch ev := event.(type) {
		case *tcell.EventResize:
			r.Sync()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				r.Clear()
				r.ShowCursor(0, 0)
				r.Fini()
				os.Exit(0)
			}

			switch ev.Key() {
			case tcell.KeyLeft:
				pos.X--
			case tcell.KeyRight:
				pos.X++
			case tcell.KeyUp:
				pos.Y--
			case tcell.KeyDown:
				pos.Y++
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
		w.SetEntityComponent(pos, e)
	}

}

func Render(r draw.Renderer, w *ecs.World) {

	for e := range w.QueryEntitiesIter(Renderable{}, Position{}) {
		pos := w.GetEntityComponent(PositionTag, e).(Position)
		renderable := w.GetEntityComponent(RenderableTag, e).(Renderable)

		r.SetContent(pos.X, pos.Y, renderable.Glyph, nil, renderable.Style)
	}
}
