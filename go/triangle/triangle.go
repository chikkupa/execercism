package triangle

import "math"

// Kind the kind of output value
type Kind int

const (
	// NaT Not a Triangle
	NaT = iota
	// Equ Equilateral triangle
	Equ
	// Iso Isosceles
	Iso
	// Sca scalene
	Sca
	// Deg Degenerate tiangle
	Deg
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || (a <= 0) || (b <= 0) || (c <= 0) || (a+b < c) || (b+c < a) || (c+a < b) {
		return NaT
	}

	if (a-b) == 0 && (b-c) == 0 {
		return Equ
	}

	if (a-b) == 0 || (b-c) == 0 || c-a == 0 {
		return Iso
	}

	if (a+b-c) == 0 || (b+c-a) == 0 || (a+c-b) == 0 {
		return Deg
	}

	return Sca
}
