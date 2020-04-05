package math32

import "math"

// Signbit reports whether x is negative or negative zero.
func Signbit(x float32) bool {
	return math.Signbit(float64(x))
}
