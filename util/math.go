package util

import "math"

// MinInt provides an implementation of math.Min for ints.
func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// MaxInt provides an implementation of math.Max for ints.
func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Distance returns the Manhattan distance between two points on a 2d grid.
// It will panic for integers > (2^53) - 1, given the size of our map we should be fine.
func Distance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))
}

// Rect is used to encapsulate 2D rectangular math
type Rect struct {
	X1, X2, Y1, Y2 int
}

// NewRect creates a new rect with the x,y coordinates of the top left corner, width and height
func NewRect(x, y, w, h int) Rect {
	return Rect{x, x + w, y, y + h}
}

// Intersect returns whether two rectangles overlap
func (r *Rect) Intersect(other Rect) bool {
	return r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y2 && r.Y2 >= other.Y1
}

// Center returns the x,y coordinates of the rectangle center point
func (r *Rect) Center() (int, int) {
	return (r.X1 + r.X2) / 2, (r.Y1 + r.Y2) / 2
}
