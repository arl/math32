package math32

import "math"

// Remainder returns the IEEE 754 floating-point remainder of x/y.
//
// Special cases are:
//	Remainder(±Inf, y) = NaN
//	Remainder(NaN, y) = NaN
//	Remainder(x, 0) = NaN
//	Remainder(x, ±Inf) = x
//	Remainder(x, NaN) = NaN
func Remainder(x, y float32) float32 {
	return float32(math.Remainder(float64(x), float64(y)))
}
