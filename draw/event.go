package draw

// ScreenEvent provides a platform-independent abstaction for input/resize events.
type ScreenEvent interface {
	ScreenEvent()
}

// ScreenEventKey provides a platform-independent abstaction for keyboard keys.
type ScreenEventKey int

const (
	KeyDummy ScreenEventKey = iota
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyControlC
	KeyRune
	KeyEscape
)

// KeyEvent represent a keyboard key press.
type KeyEvent struct {
	Rune rune
	Key  ScreenEventKey
}

func (*KeyEvent) ScreenEvent() {}

// MouseButton provides a platform-independent abstaction for mouse buttons.
type MouseButton int

const (
	Primary MouseButton = iota
	Secondary
)

// MouseEvent represents mousemoves/clicks.
type MouseEvent struct {
	X, Y   int
	Button MouseButton
}

func (*MouseEvent) ScreenEvent() {}

// ResizeEvent represents render surface resizing.
type ResizeEvent struct {
	Width, Height int
}

func (*ResizeEvent) ScreenEvent() {}
