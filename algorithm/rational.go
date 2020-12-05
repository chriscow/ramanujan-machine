package algorithm

import (
	"fmt"
	"math"
)

type RationalFunc struct {
	A, B float64
}

// Solve is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func (e RationalFunc) Solve() (float64, error) {
	if e.B == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: b == 0")
	}

	return e.A / e.B, nil
}
