package systems

import (
	"math/rand"

	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/resources"
)

// RandomPlayer is a system used for testing.
// By registering this, it is possible to simulate a game without player input
func RandomPlayer(w *ecs.World) {
	r := ecs.GetResource[resources.Renderer](w)
	state := ecs.GetResource[resources.GameState](w)
	if state != resources.PlayerTurn {
		return
	}
	switch rand.Intn(4) {
	case 0:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyLeft, Rune: ' '})
	case 1:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyRight, Rune: ' '})
	case 2:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyUp, Rune: ' '})
	case 3:
		r.PostEvent(&draw.KeyEvent{Key: draw.KeyDown, Rune: ' '})
	}
}
