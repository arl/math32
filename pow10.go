package math32

import "math"

// Pow10 returns 10**n, the base-10 exponential of n.
//
// Special cases are:
//	Pow10(n) =    0 for n < -323
//	Pow10(n) = +Inf for n > 308
func Pow10(e int) float32 {
	return float32(math.Pow10(e))
}
