package algorithm

import (
	"fmt"
	"log"
	"math"
	"ramanujan/sequence"
	"ramanujan/slice"

	"github.com/shamaton/msgpack"
)

// NestedRadical struct supports serialization of generator config
type NestedRadical struct {
	A, B sequence.Generator
}

// GetType returns the type ID representing this algorithm
func (nr NestedRadical) GetType() AlgoType {
	return NestedRad
}

// Solve calculates a value by using the sequence of values passed in a
// nested radical of the form:
//
//		sqrt(a + b * sqrt(a + b * sqrt(a + b * sqrt([ ... ]))))
//
// See: https://www.johndcook.com/blog/2013/09/13/ramanujans-nested-radical/
//
// arguments a and b are expected to be the same length
func (nr NestedRadical) Solve() <-chan Solution {
	ch := make(chan Solution, 100)

	go func() {
		defer close(ch)
		for a := range nr.A.Next() {
			for b := range nr.B.Next() {
				res, err := nr.solve(a, b)
				if err != nil {
					log.Fatal("continued fraction misconfigured", err)
				}

				arg := [][]float64{a, b}
				b, err := msgpack.Encode(arg)
				if err != nil {
					panic(err)
				}

				ch <- Solution{Type: NestedRad, Args: b, Result: res}
			}
		}
	}()

	return ch
}

func (nr NestedRadical) solve(a, b []float64) (float64, error) {
	var root float64
	root = 1

	if len(a) != len(b) {
		return math.NaN(), fmt.Errorf("invalid argument: arrays must be same length")
	}

	slice.Reverse(a)
	slice.Reverse(b)

	for i := range a {
		root = math.Sqrt(b[i]*root + a[i])
	}

	return root, nil
}
