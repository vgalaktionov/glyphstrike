package systems

import (
	"log"
	"os"
	"time"

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

	playerEnt, err := ecs.QueryEntitiesSingle(w, Player{}, Position{})
	if err != nil {
		log.Fatal("no player found")
	}

	switch ev := event.(type) {
	case *draw.ResizeEvent:
		r.Sync()

	case *draw.KeyEvent:
		if ev.Key == draw.KeyEscape || ev.Key == draw.KeyControlC {
			r.CleanUp()
			os.Exit(0)
		}

		chord := maybeChord(r)

		var deltaX, deltaY int

		switch true {
		case isLeft(ev) && isUp(chord), isUp(ev) && isLeft(chord):
			deltaX--
			deltaY--
		case isRight(ev) && isUp(chord), isUp(ev) && isRight(chord):
			deltaX++
			deltaY--
		case isLeft(ev) && isDown(chord), isDown(ev) && isLeft(chord):
			deltaX--
			deltaY++
		case isRight(ev) && isDown(chord), isDown(ev) && isRight(chord):
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

		default:
			log.Printf("Unbound key pressed: %b", ev.Rune)
		}

		tryMovePlayer(w, playerEnt, deltaX, deltaY)

	case *draw.MouseEvent:
		return
	}

}

// tryMovePlayer handles player movement and attack input (via bump to fight)
func tryMovePlayer(w *ecs.World, playerEnt ecs.Entity, deltaX, deltaY int) {
	playerPos := ecs.MustGetEntityComponent[Position](w, playerEnt)
	destX := playerPos.X + deltaX
	destY := playerPos.Y + deltaY

	m := ecs.GetResource[*resources.Map](w)

	for _, potentialTarget := range m.TileContents[destX][destY] {
		if ecs.HasEntityComponent[CombatStats](w, ecs.Entity(potentialTarget)) {
			ecs.SetEntityComponent(w, WantsToMelee{Target: potentialTarget}, playerEnt)
			return
		}
	}

	if !m.BlockedTiles[destX][destY] {
		ecs.SetEntityComponent(w, Position{X: destX, Y: destY}, playerEnt)
	}
}

// isUp checks whether a key event represents a valid "up" key press.
func isUp(keyEv *draw.KeyEvent) bool {
	if keyEv == nil {
		return false
	}
	return draw.KeyUp == keyEv.Key || (draw.KeyRune == keyEv.Key && keyEv.Rune == 'w')
}

// isDown checks whether a key event represents a valid "down" key press.
func isDown(keyEv *draw.KeyEvent) bool {
	if keyEv == nil {
		return false
	}
	return draw.KeyDown == keyEv.Key || (draw.KeyRune == keyEv.Key && keyEv.Rune == 's')
}

// isLeft checks whether a key event represents a valid "left" key press.
func isLeft(keyEv *draw.KeyEvent) bool {
	if keyEv == nil {
		return false
	}
	return draw.KeyLeft == keyEv.Key || (draw.KeyRune == keyEv.Key && keyEv.Rune == 'a')
}

// isRight checks whether a key event represents a valid "right" key press.
func isRight(keyEv *draw.KeyEvent) bool {
	if keyEv == nil {
		return false
	}
	return draw.KeyRight == keyEv.Key || (draw.KeyRune == keyEv.Key && keyEv.Rune == 'd')
}

const ChordMS = 50

// maybeChord is used to work around the limitation of the terminal only receiving a single keypress at a time.
// It returns the next key pressed within `ChordMS`, or nil.
func maybeChord(r draw.Screen) *draw.KeyEvent {
	ev := make(chan draw.ScreenEvent)
	timer := time.NewTimer(ChordMS * time.Millisecond)

	go func() {
		defer close(ev)
		ev <- r.PollEvent()
	}()

	for {
		select {
		case event := <-ev:
			switch keyEv := event.(type) {
			case *draw.KeyEvent:
				if keyEv.Key == draw.KeyDummy {
					return nil
				}
				return keyEv
			default:
				return nil
			}
		case <-timer.C:
			// send dummy key event
			r.PostEvent(&draw.KeyEvent{})
		}
	}
}
