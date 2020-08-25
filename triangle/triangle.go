// Package triangle implements a function to qualify a triangle
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides takes the lenghts of the 3 sides of a triangle as input
// and outputs the triangle kind.
func KindFromSides(a, b, c float64) Kind {
	// Exclude Nan, Inf, null and negative numbers
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) ||
		a <= 0 || b <= 0 || c <= 0 ||
		a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// All valid cases
	if a == b {
		if b == c {
			// Only possible Equ case
			return Equ
		}
		return Iso
	}
	if a == c || b == c {
		return Iso
	}
	return Sca
}
