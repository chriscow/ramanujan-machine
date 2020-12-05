package algorithm

import (
	"fmt"
	"math"
	"ramanujan/utils"
)

type NestedRadical struct {
	A, B []float64
}

// Solve calculates a value by using the sequence of values passed in a
// nested radical of the form:
//
//		sqrt(a + b * sqrt(a + b * sqrt(a + b * sqrt([ ... ]))))
//
// See: https://www.johndcook.com/blog/2013/09/13/ramanujans-nested-radical/
//
// arguments a and b are expected to be the same length
func (nr NestedRadical) Solve() (float64, error) {
	var root float64
	root = 1

	a, b := nr.A, nr.B

	if len(a) != len(b) {
		return math.NaN(), fmt.Errorf("invalid argument: arrays must be same length")
	}

	utils.Reverse(a)
	utils.Reverse(b)

	for i := range a {
		root = math.Sqrt(b[i]*root + a[i])
	}

	return root, nil
}
