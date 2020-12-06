package algorithm

import (
	"fmt"
	"log"
	"math"
	"ramanujan/sequence"
)

// RationalFunc supports serialization of its generators
type RationalFunc struct {
	A, B sequence.Generator
}

// Solve is simply a rational function that divides a by b
// The two main algorithms are the Rational for the left hand side
// and ContinuedFraction for the right hand side.
func (e RationalFunc) Solve() <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)
		for a := range e.A.Next() {
			for b := range e.B.Next() {
				for i := range a {
					for j := range b {
						res, err := e.solve(a[i], b[j])
						if err != nil {
							log.Fatal(err)
						}
						ch <- res
					}
				}
			}
		}
	}()

	return ch
}

func (e RationalFunc) solve(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: b == 0")
	}

	return a / b, nil
}
