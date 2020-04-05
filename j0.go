package math32

import "math"

// J0 returns the order-zero Bessel function of the first kind.
//
// Special cases are:
//	J0(±Inf) = 0
//	J0(0) = 1
//	J0(NaN) = NaN
func J0(x float32) float32 {
	return float32(math.J0(float64(x)))
}

// Y0 returns the order-zero Bessel function of the second kind.
//
// Special cases are:
//	Y0(+Inf) = 0
//	Y0(0) = -Inf
//	Y0(x < 0) = NaN
//	Y0(NaN) = NaN
func Y0(x float32) float32 {
	return float32(math.Y0(float64(x)))
}
