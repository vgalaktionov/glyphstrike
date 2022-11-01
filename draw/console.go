//go:build !js
// +build !js

package draw

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

type ConsoleRenderer struct {
	tcell.Screen
}

func NewScreen() Screen {
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(DEFAULT_STYLE)
	screen.EnableMouse()
	screen.Clear()
	return &ConsoleRenderer{screen}
}

func (cr *ConsoleRenderer) CleanUp() {
	cr.Clear()
	cr.ShowCursor(0, 0)
	cr.Fini()
}

func (cr *ConsoleRenderer) PollEvent() ScreenEvent {
	ev := cr.Screen.PollEvent()
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyCtrlC:
			return &KeyEvent{Key: KeyControlC, Rune: 'c'}
		case tcell.KeyEscape:
			return &KeyEvent{Key: KeyEscape, Rune: ' '}
		case tcell.KeyF64:
			return &KeyEvent{Key: KeyDummy, Rune: ' '}
		case tcell.KeyUp:
			return &KeyEvent{Key: KeyUp, Rune: ' '}
		case tcell.KeyDown:
			return &KeyEvent{Key: KeyDown, Rune: ' '}
		case tcell.KeyLeft:
			return &KeyEvent{Key: KeyLeft, Rune: ' '}
		case tcell.KeyRight:
			return &KeyEvent{Key: KeyRight, Rune: ' '}
		default:
			return &KeyEvent{Key: KeyRune, Rune: event.Rune()}
		}
	case *tcell.EventMouse:
		return &MouseEvent{}

	case *tcell.EventResize:
		return &ResizeEvent{}
	default:
		log.Panicln("unknown key")
		return nil
	}
}

func (cr *ConsoleRenderer) PostEvent(ev ScreenEvent) error {
	switch event := ev.(type) {
	case *KeyEvent:
		switch event.Key {
		case KeyUp:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone))
		case KeyLeft:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyLeft, ' ', tcell.ModNone))
		case KeyRight:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyRight, ' ', tcell.ModNone))
		case KeyDown:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone))
		case KeyDummy:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyF64, ' ', tcell.ModNone))
		case KeyEscape:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyEscape, ' ', tcell.ModNone))
		case KeyRune:
			return cr.Screen.PostEvent(tcell.NewEventKey(tcell.KeyRune, event.Rune, tcell.ModNone))
		}
	case *MouseEvent:
		switch event.Button {
		case Primary:
			return cr.Screen.PostEvent(tcell.NewEventMouse(event.X, event.Y, tcell.ButtonPrimary, tcell.ModNone))
		case Secondary:
			return cr.Screen.PostEvent(tcell.NewEventMouse(event.X, event.Y, tcell.ButtonSecondary, tcell.ModNone))
		}

	case *ResizeEvent:
		return cr.Screen.PostEvent(tcell.NewEventResize(event.Width, event.Height))
	default:
		return nil
	}
	return fmt.Errorf("invalid event: %s", ev)
}

func (cr ConsoleRenderer) SetCellContent(x int, y int, primary rune, foreground, background ColorName) {
	cr.Screen.SetContent(x, y, primary, nil, palette[foreground][background])
}

// palette is a package-private lookup of rendering styles, precomputed on init
var palette [len(Colors)][len(Colors)]tcell.Style

func init() {
	for i, fg := range Colors {
		for j, bg := range Colors {
			palette[i][j] = DEFAULT_STYLE.Foreground(tcell.NewHexColor(int32(fg)).TrueColor()).Background(tcell.NewHexColor(int32(bg)).TrueColor())
		}
	}
}

var DEFAULT_STYLE tcell.Style = tcell.StyleDefault.Background(tcell.ColorBlack.TrueColor()).Foreground(tcell.ColorWhite.TrueColor())
