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
//
// Generators create arrays of floats so we are only taking the first value
// in that array and this is expected. This method will cause a fatal exit if a generator
// returns more than a single element.
func (e RationalFunc) Solve() <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)
		for a := range e.A.Next() {
			if len(a) > 1 {
				log.Fatal("RationalFunc generator A misconfigured to return more than a single value")
			}
			for b := range e.B.Next() {
				if len(b) > 1 {
					log.Fatal("RationalFunc generator B misconfigured to return more than a single value")
				}
				res, err := e.solve(a[0], b[0])
				if err != nil {
					log.Fatal("rational function generator misconfigured", err)
				}
				ch <- res
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
