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

func PlayerInput(r draw.Renderer, w *ecs.World) {
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
			}

			var deltaX, deltaY int
			switch ev.Key() {
			case tcell.KeyLeft:
				deltaX--
			case tcell.KeyRight:
				deltaX++
			case tcell.KeyUp:
				deltaY--
			case tcell.KeyDown:
				deltaY++
			}

			destX := playerPos.X + deltaX
			destY := playerPos.Y + deltaY

			// Walls are solid
			m := w.GetResource(resources.MapTag).(resources.Map)
			if m.Tiles[destX][destY] != resources.WallTile {
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
