package math32

import "math"

// Nextafter32 returns the next representable float32 value after x towards y.
//
// Special cases are:
//	Nextafter(x, x)   = x
//	Nextafter(NaN, y) = NaN
//	Nextafter(x, NaN) = NaN
func Nextafter(x, y float32) (r float32) {
	return math.Nextafter32(x, y)
}
