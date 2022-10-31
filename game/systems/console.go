package systems

import (
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
			draw.DrawStr(r, 1, y, draw.ColorFromPalette(draw.White, draw.Black), line[:util.MinInt(UIOffsetX-1, len(line))])
		}
	}
}
