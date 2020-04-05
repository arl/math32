package math32

import "math"

// J1 returns the order-one Bessel function of the first kind.
//
// Special cases are:
//	J1(±Inf) = 0
//	J1(NaN) = NaN
func J1(x float32) float32 {
	return float32(math.J1(float64(x)))
}

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
//	Y1(+Inf) = 0
//	Y1(0) = -Inf
//	Y1(x < 0) = NaN
//	Y1(NaN) = NaN
func Y1(x float32) float32 {
	return float32(math.Y1(float64(x)))
}
