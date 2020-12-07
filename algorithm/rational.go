package algorithm

import (
	"fmt"
	"log"
	"math"
	"ramanujan/sequence"

	"github.com/shamaton/msgpack"
)

// RationalFunc supports serialization of its generators
type RationalFunc struct {
	A, B sequence.Generator
}

// GetType returns the type ID representing this algorithm
func (rf RationalFunc) GetType() AlgoType {
	return Rational
}

// Solve is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func (rf RationalFunc) Solve() <-chan Solution {
	ch := make(chan Solution)

	go func() {
		defer close(ch)
		for a := range rf.A.Next() {
			for b := range rf.B.Next() {
				for i := range a {
					for j := range b {
						res, err := rf.solve(a[i], b[j])
						if err != nil {
							log.Fatal(err)
						}

						arg := []float64{a[i], b[i]}
						b, err := msgpack.Encode(arg)
						if err != nil {
							panic(err)
						}

						ch <- Solution{Type: Rational, Args: b, Result: res}
					}
				}
			}
		}
	}()

	return ch
}

func (rf RationalFunc) solve(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: b == 0")
	}

	return a / b, nil
}
