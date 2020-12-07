package main

import (
	"ramanujan/algorithm"
	algo "ramanujan/algorithm"
	"ramanujan/sequence"
)

// side contains all the information to generate a series of calculated values
// from a list of algorithms
type side struct {
	Algorithms []algo.Solver
	PostProc   bool
	Ignore     []float64
}

// Solve takes an equation's side configuration (recall that an equation has a
// left and a right hand side) and loops over each algorithm that will calculate
// the values for that side.
//
// An algorithm takes two sequences. Each sequence has a generator that generates
// the values for the sequence, so loop over the array of all input sequence
// generators and use them to generate a series of sequences.
//
// FINALLY, execute the algorithm with each sequence generated. Essentially this
// is calculating the algorithm value with a brute force approach with every
// possible set of values within the sequence range generated.
func (conf side) Solve() <-chan algorithm.Solution {
	ch := make(chan algorithm.Solution)

	go func() {
		defer close(ch)

		for _, algFunc := range conf.Algorithms {

			for sln := range algFunc.Solve() {
				// see if we are supposed to ignore the returned value
				ignore := false
				for _, i := range conf.Ignore {
					if sln.Result == i {
						ignore = true
						break
					}
				}

				if !ignore {
					ch <- sln
				}
			}
		}
	}()

	return ch
}

//__[ Sample Equation Sides ]___________________________________________________
//
// These sample side configurations are great for testing
//
func lhsFindsConstants(constants []float64) side {
	ones := make([]float64, len(constants))
	for i := range ones {
		ones[i] = 1
	}

	a := sequence.Constant{Values: constants}
	b := sequence.Constant{Values: ones}

	rf := algorithm.RationalFunc{A: a, B: b}

	solvers := []algorithm.Solver{rf}

	return side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

// Configuration that finds e
func rhsFindsE() side {
	a := sequence.Polynomial{
		A:    []float64{3, 4},
		B:    []float64{1, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 1},
		B:    []float64{-1, 0},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	cf := algorithm.ContinuedFraction{A: a, B: b}
	solvers := []algorithm.Solver{cf}

	return side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

func rhsPhiandE() side {
	a := sequence.Polynomial{
		A:    []float64{1, 4},
		B:    []float64{0, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 2},
		B:    []float64{-1, 1},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	cf := algorithm.ContinuedFraction{A: a, B: b}
	solvers := []algorithm.Solver{cf}

	return side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

// Configuration that finds phi for both contined fraction and nested radical
func rhsPhiCFandNR() side {
	a := sequence.Polynomial{
		A:    []float64{1, 4},
		B:    []float64{0, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 2},
		B:    []float64{-1, 1},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	cf := algorithm.ContinuedFraction{A: a, B: b}
	nr := algorithm.NestedRadical{A: a, B: b}
	solvers := []algorithm.Solver{cf, nr}

	return side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

func rhsFindsSqrt3EandPhi() side {
	a1 := sequence.Integer{
		Digits:   []int{1, 2},
		Count:    2,
		Repeat:   100,
		Prefix:   []int{1},
		PfxCount: 1,
	}

	a2 := sequence.Polynomial{
		A:    []float64{1, 4},
		B:    []float64{0, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 2},
		B:    []float64{-1, 1},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	cf1 := algorithm.ContinuedFraction{A: a1, B: b}
	cf2 := algorithm.ContinuedFraction{A: a2, B: b}
	nr1 := algorithm.NestedRadical{A: a1, B: b}
	nr2 := algorithm.NestedRadical{A: a2, B: b}

	solvers := []algorithm.Solver{cf1, cf2, nr1, nr2}

	return side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}
