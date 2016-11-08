package math32

import "math"

var (
	Epsilon32 float32
	NaN       float32
)

func init() {
	Epsilon32 = math.Nextafter32(1, 2) - 1
	NaN = float32(math.NaN())
}

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

// Max returns the larger of x or y.
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func Max(x, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}

// Min returns the smaller of x or y.
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func Min(x, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim(x, y float32) float32 {
	return float32(math.Dim(float64(x), float64(y)))
}

// Approx returns true if x ~= y
func Approx(x, y float32) bool {
	eps := Epsilon32 * 100
	return Abs(x-y) < eps*(1.0+Max(Abs(x), Abs(y)))
}
