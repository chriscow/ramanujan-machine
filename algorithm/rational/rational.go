package rational

import (
	"math"
)

// Rational is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func Solve(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}

	return a / b
}
