package draw

// RGB is a color triplet.
type RGB struct {
	R, G, B uint8
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
	Black
	Red
	Yellow
	BlueGreen
)

var Colors = [9]int{
	0xECEFF4,
	0xD8DEE9,
	0x4C566A,
	0x3B4252,
	0x2E3440,
	0xBF616A,
	0xEBCB8B,
	0x8FBCBB,
}

