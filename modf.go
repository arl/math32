package math32

import "math"

// Modf returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
//	Modf(±Inf) = ±Inf, NaN
//	Modf(NaN) = NaN, NaN
func Modf(f float32) (int float32, frac float32) {
	i, f64 := math.Modf(float64(f))
	return float32(i), float32(f64)
}
