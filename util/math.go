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
