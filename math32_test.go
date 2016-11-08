package math32

import "testing"

func TestApprox(t *testing.T) {
	approxTests := []struct {
		x, y float32
		want bool // true means equal
	}{
		{1.0, 1.0, true},
		{1.0, 1.000001, true},
		{1.0, 1.00001, true},
		{1.0, 1.0001, false},
		{1.0, 1.001, false},
		{1.0, 1.01, false},
		{1.0, 0.999999, true},
		{1.0, 0.99999, true},
		{1.0, 0.9999, false},
		{1.0, 0.999, false},
		{1.0, 0.99, false},
		{0.0, 0.000001, true},
		{0.0, 0.00001, true},
		{0.0, 0.0001, false},
		{0.0, 0.001, false},
		{0.0, 0.01, false},
		{0.0, -0.000001, true},
		{0.0, -0.00001, true},
		{0.0, -0.0001, false},
		{0.0, -0.001, false},
		{0.0, -0.01, false},
		{1e12, 1e12 + 0.000001, true},
		{1e12, 1e12 + 0.00001, true},
		{1e12, 1e12 + 0.0001, true},
		{1e12, 1e12 + 0.001, true},
		{1e12, 1e12 + 0.01, true},
		{1e12, 1e12 - 0.000001, true},
		{1e12, 1e12 - 0.00001, true},
		{1e12, 1e12 - 0.0001, true},
		{1e12, 1e12 - 0.001, true},
		{1e12, 1e12 - 0.01, true},
		{NaN, 0, false},
		{NaN, NaN, false},
	}

	for _, tt := range approxTests {
		got := Approx(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("Approx(%f, %f) == %t, want %t", tt.x, tt.y, got, tt.want)
		}
	}
}
