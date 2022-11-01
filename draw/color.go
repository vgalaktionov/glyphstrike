package draw

import "github.com/gdamore/tcell/v2"

// RGB is a color triplet.
type RGB struct {
	R, G, B int32
}

// Color provides a platform-independent color abstraction.
type Color struct {
	Foreground RGB
	Background RGB
}

// ColorName is a color in our palette
type ColorName int

const (
	White ColorName = iota
	LightGray
	DarkGray
	DarkerGray
	BlueGreen
	Black
	Red
	Green
	Yellow
	Transparent
)

// ColorFromPalette resolves color shortcuts by name
func ColorFromPalette(name1, name2 ColorName) Color {
	colors := []RGB{{0, 0, 0}, {0, 0, 0}}
	for i, name := range []ColorName{name1, name2} {
		switch name {
		case White:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0xECEFF4).RGB()
		case LightGray:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0xD8DEE9).RGB()
		case DarkGray:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0x4C566A).RGB()
		case DarkerGray:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0x3B4252).RGB()
		case Black:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0x2E3440).RGB()
		case BlueGreen:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0x8FBCBB).RGB()
		case Red:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0xBF616A).RGB()
		case Yellow:
			colors[i].R, colors[i].G, colors[i].B = tcell.NewHexColor(0xEBCB8B).RGB()
		}
	}
	return Color{colors[0], colors[1]}
}
