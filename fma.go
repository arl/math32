package math32

import "math"

// FMA returns x * y + z, computed with only one rounding.
// (That is, FMA returns the fused multiply-add of x, y, and z.)
func FMA(x, y, z float32) float32 {
	return float32(math.FMA(float64(x), float64(y), float64(z)))
}
