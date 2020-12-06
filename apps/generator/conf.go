package main

import (
	"ramanujan/algorithm"
	"ramanujan/sequence"
)

const (
	defaultArgs = "default-args.json"
	argsUsage   = "-args args.json"
)

func checkEnv() error {

	// TODO: verify required environment variables

	return nil
}

func loadArgs(argsFile string) error {

	// setDefaults()

	return nil
}

//__[ Sample Configurations ]___________________________________________________
//
// The following functions will configure side configurations for interesting
// results.  It will either find some well known constants, good for testing or
// or they may define runs of a certain size for scale / load testing.
//
func tiny() (lhs, rhs Side) {

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	rhs = rhsFindsE()
	lhs = lhsFindsConstants()

	return
}

func lhsFindsConstants() Side {
	rf := algorithm.RationalFunc{
		A: sequence.Polynomial{
			A: []float64{0, 1}, // args are reversed!
			B: []float64{1, 2}, // 0 + 1x + 0x^2 where x == the const
			C: []float64{0, 1},
		},
		B: sequence.Polynomial{ // don't forget the args are reversed
			A: []float64{1, 2}, // we are just dividing by 1
			B: []float64{0, 1}, // 1 + 0x + 0x^2
			C: []float64{0, 1}, // A +  B + C
		},
	}

	solvers := []algorithm.Solver{rf}

	return Side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

func rhsPhiandE() Side {
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

	return Side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

// Configuration that finds phi for both contined fraction and nested radical
func rhsPhiCFandNR() Side {
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

	return Side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

// Configuration that finds e
func rhsFindsE() Side {
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

	return Side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}
}

// Configuration that generates just enough range to find BOTH phi and e
// func findsEandPhi() Side {
// 	return Side{
// 		algorithms: []algorithm.Solver{
// 			contfrac.Solve,
// 		},
// 		postproc:  false,
// 		blacklist: []float64{-2, -1, 0, 1, 2},
// 		aSeqs: []seq.Config{
// 			seq.Config{
// 				Generator: sequence.Polynomial,
// 				Coeff: [][]float64{
// 					[]float64{1, 4},
// 					[]float64{0, 2},
// 					[]float64{0, 1},
// 				},
// 				RangeLow:  0,
// 				RangeHigh: 200,
// 			},
// 		},
// 		bSeqs: []seq.Config{
// 			seq.Config{
// 				Generator: sequence.Polynomial,
// 				Coeff: [][]float64{
// 					[]float64{0, 2},
// 					[]float64{-1, 1},
// 					[]float64{0, 1},
// 				},
// 				RangeLow:  0,
// 				RangeHigh: 200,
// 			},
// 		},
// 	}
// }
