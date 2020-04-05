package math32

import "math"

// Erfinv returns the inverse error function of x.
//
// Special cases are:
//	Erfinv(1) = +Inf
//	Erfinv(-1) = -Inf
//	Erfinv(x) = NaN if x < -1 or x > 1
//	Erfinv(NaN) = NaN
func Erfinv(x float32) float32 {
	return float32(math.Erfinv(float64(x)))
}
