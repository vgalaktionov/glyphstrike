package util

import "math"

// MinInt provides an implementation of math.Min for ints. It is only safe for integers fitting a float64.
func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// MaxInt provides an implementation of math.Max for ints. It is only safe for integers fitting a float64.
func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// Rect 
type Rect struct {
	X1, X2, Y1, Y2 int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{x, x + w, y, y + h}
}

func (r *Rect) Intersect(other Rect) bool {
	return r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y2 && r.Y2 >= other.Y1
}

func (r *Rect) Center() (int, int) {
	return (r.X1 + r.X2) / 2, (r.Y1 + r.Y2) / 2
}
