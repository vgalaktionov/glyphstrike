package systems

import (
	"fmt"

	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/game/events"
	. "github.com/vgalaktionov/roguelike-go/game/resources"
	"github.com/vgalaktionov/roguelike-go/util"
)

// Console system runs in separate goroutine (as eventsystem) and processes ConsoleEvent messages,
// writing the last n messages to the screen.
func Console(w *ecs.World) {
	ch := ecs.GetEventChannel[events.ConsoleEvent](w)
	lines := []string{}
	r := ecs.GetResource[Renderer](w)

	_, maxY := r.Size()
	// loop forever, as we run in background
	for {

		ev := <-ch
		lines = append(lines, ev.(events.ConsoleEvent).Message)

		for y, line := range lines[util.MaxInt(len(lines)-maxY, 0):] {
			// We want to repaint the entire line, so pad it out
			str := fmt.Sprintf("%-*s", UIOffsetX-1, line)
			draw.DrawStr(r, 1, y+UIOffsetY+1, draw.White, draw.Black, str)
		}
	}
}
