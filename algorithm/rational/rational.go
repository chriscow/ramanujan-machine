package rational

import (
	"fmt"
	"math"
)

// Solve is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func Solve(a, b []float64) (float64, error) {
	if len(a) != 1 || len(b) != 1 {
		return math.NaN(), fmt.Errorf("invalid argument: length of a and b must be 1")
	}

	if b[0] == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: b == 0")
	}

	return a[0] / b[0], nil
}
