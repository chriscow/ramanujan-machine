package rational

import (
	"fmt"
	"math"
)

// Rational is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func Solve(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: b == 0")
	}

	return a / b, nil
}
