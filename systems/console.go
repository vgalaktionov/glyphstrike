package systems

import (
	"github.com/gdamore/tcell/v2"
	"github.com/vgalaktionov/roguelike-go/draw"
	"github.com/vgalaktionov/roguelike-go/ecs"
	"github.com/vgalaktionov/roguelike-go/events"
	"github.com/vgalaktionov/roguelike-go/util"
)

type consoleBuffer struct {
	lines []string
	dirty bool
}

func Console(r draw.Renderer, w *ecs.World) {
	ch := w.GetEventChannel(events.ConsoleEventTag)
	cb := consoleBuffer{}

	_, maxY := r.Size()
	for {

		for len(ch) > 1 {
			ev := <-ch
			cb.lines = append(cb.lines, ev.(events.ConsoleEvent).Message)
			cb.dirty = true

		}
		if cb.dirty {
			for y, line := range cb.lines[util.MaxInt(len(cb.lines)-maxY, 0):] {
				draw.DrawStr(r, 1, y, tcell.StyleDefault, line[:util.MinInt(UIOffsetX-1, len(line))])
			}
			cb.dirty = false
		}
	}
}
