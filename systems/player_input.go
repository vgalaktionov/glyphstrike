package systems

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/resources"

	//lint:ignore ST1001 dot importing components makes it much more readable in this case
	. "github.com/vgalaktionov/roguelike-go/components"
)

func HandlePlayerInput(r draw.Renderer, w *ecs.World) {
	event := r.PollEvent()
	for e := range w.QueryEntitiesIter(Player{}, Position{}) {
		playerPos := w.GetEntityComponent(PositionTag, e).(Position)

		switch ev := event.(type) {
		case *tcell.EventResize:
			r.Sync()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				r.Clear()
				r.ShowCursor(0, 0)
				r.Fini()
				os.Exit(0)
				break
			}

			var deltaX, deltaY int

			switch true {
			case tcell.KeyLeft == ev.Key():
				fallthrough
			case tcell.KeyRune == ev.Key() && ev.Rune() == 'h':
				deltaX--
			case tcell.KeyRight == ev.Key():
				fallthrough
			case tcell.KeyRune == ev.Key() && ev.Rune() == 'l':
				deltaX++
			case tcell.KeyUp == ev.Key():
				fallthrough
			case tcell.KeyRune == ev.Key() && ev.Rune() == 'k':
				deltaY--
			case tcell.KeyDown == ev.Key():
				fallthrough
			case tcell.KeyRune == ev.Key() && ev.Rune() == 'j':
				deltaY++
			}

			destX := playerPos.X + deltaX
			destY := playerPos.Y + deltaY

			// Walls are solid
			m := w.GetResource(resources.MapTag).(*resources.Map)
			if !m.IsOpaque(destX, destY) {
				w.SetEntityComponent(Position{X: destX, Y: destY}, e)
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

}
