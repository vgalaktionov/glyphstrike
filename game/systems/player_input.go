package systems

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	. "github.com/vgalaktionov/roguelike-go/game/components"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// HandlePlayerInput processes keyboard/mouse input and resize events.
func HandlePlayerInput(w *ecs.World) {
	gs := ecs.GetResource[resources.GameState](w)
	if gs != resources.PlayerTurn {
		return
	}
	r := ecs.GetResource[resources.Renderer](w)

	event := r.PollEvent()

	for e := range ecs.QueryEntitiesIter(w, Player{}, Position{}) {
		playerPos := ecs.GetEntityComponent[Position](w, e)

		switch ev := event.(type) {
		case *tcell.EventResize:
			r.Sync()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				r.CleanUp()
				os.Exit(0)
			}

			chord := maybeChord(r)

			var deltaX, deltaY int

			switch true {
			case ev.Key() == tcell.KeyUpLeft, isLeft(ev) && isUp(chord), isUp(ev) && isLeft(chord):
				deltaX--
				deltaY--
			case ev.Key() == tcell.KeyUpRight, isRight(ev) && isUp(chord), isUp(ev) && isRight(chord):
				deltaX++
				deltaY--
			case ev.Key() == tcell.KeyDownLeft, isLeft(ev) && isDown(chord), isDown(ev) && isLeft(chord):
				deltaX--
				deltaY++
			case ev.Key() == tcell.KeyDownRight, isRight(ev) && isDown(chord), isDown(ev) && isRight(chord):
				deltaX++
				deltaY++
			case isLeft(ev):
				deltaX--
			case isRight(ev):
				deltaX++
			case isUp(ev):
				deltaY--
			case isDown(ev):
				deltaY++

			case tcell.KeyUpLeft == ev.Key(), tcell.KeyUpRight == ev.Key(), tcell.KeyDownRight == ev.Key(), tcell.KeyDownLeft == ev.Key():
				log.Print("Diagonal key pressed")
			default:
				log.Printf("Unbound key pressed: %b", ev.Rune())
			}

			destX := playerPos.X + deltaX
			destY := playerPos.Y + deltaY

			m := ecs.GetResource[resources.Map](w)
			if !m.BlockedTiles[destX][destY] {
				ecs.SetEntityComponent(w, Position{X: destX, Y: destY}, e)
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

// isUp checks whether a key event represents a valid "up" key press.
func isUp(keyEv *tcell.EventKey) bool {
	if keyEv == nil {
		return false
	}
	return tcell.KeyUp == keyEv.Key() || (tcell.KeyRune == keyEv.Key() && keyEv.Rune() == 'w')
}

// isDown checks whether a key event represents a valid "down" key press.
func isDown(keyEv *tcell.EventKey) bool {
	if keyEv == nil {
		return false
	}
	return tcell.KeyDown == keyEv.Key() || (tcell.KeyRune == keyEv.Key() && keyEv.Rune() == 's')
}

// isLeft checks whether a key event represents a valid "left" key press.
func isLeft(keyEv *tcell.EventKey) bool {
	if keyEv == nil {
		return false
	}
	return tcell.KeyLeft == keyEv.Key() || (tcell.KeyRune == keyEv.Key() && keyEv.Rune() == 'a')
}

// isRight checks whether a key event represents a valid "right" key press.
func isRight(keyEv *tcell.EventKey) bool {
	if keyEv == nil {
		return false
	}
	return tcell.KeyRight == keyEv.Key() || (tcell.KeyRune == keyEv.Key() && keyEv.Rune() == 'd')
}

const ChordMS = 50

// maybeChord is used to work around the limitation of the terminal only receiving a single keypress at a time.
// It returns the next key pressed within `ChordMS`, or nil.
func maybeChord(r draw.Screen) *tcell.EventKey {
	ev := make(chan tcell.Event)
	timer := time.NewTimer(ChordMS * time.Millisecond)

	go func() {
		defer close(ev)
		ev <- r.PollEvent()
	}()

	for {
		select {
		case event := <-ev:
			switch keyEv := event.(type) {
			case *tcell.EventKey:
				if keyEv.Key() == tcell.KeyF64 {
					return nil
				}
				return keyEv
			default:
				return nil
			}
		case <-timer.C:
			// send dummy key event
			r.PostEvent(tcell.NewEventKey(tcell.KeyF64, ' ', tcell.ModNone))
		}
	}
}
