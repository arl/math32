package math32

import "math"

var (
	Epsilon32 float32
)

func init() {
	Epsilon32 = math.Nextafter32(1, 2) - 1
}

// Approx returns true if x ~= y
func Approx(x, y float32) bool {
	eps := Epsilon32 * 100
	return Abs(x-y) < eps*(1.0+Max(Abs(x), Abs(y)))
}
