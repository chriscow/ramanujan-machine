package main

import (
	"math"
	"ramanujan/algorithm"
	"ramanujan/sequence"
)

//__[ Sample Configurations ]___________________________________________________
//
// The following create side configurations for interesting
// results.  It will either find some well known constants, good for testing or
// or they may define runs of a certain size for scale / load testing.
//
func tinyConf() appConf {
	constants := []float64{math.E}

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	lhs := lhsFindsConstants(constants)
	rhs := rhsFindsE()

	return appConf{
		Constants: constants,
		LHS:       lhs,
		RHS:       rhs,
	}
}

func smallConf() appConf {
	constants := []float64{math.E, math.Phi, math.Sqrt(3)}

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	lhs := lhsFindsConstants(constants)
	rhs := rhsFindsSqrt3EandPhi()

	return appConf{
		Constants: constants,
		LHS:       lhs,
		RHS:       rhs,
	}
}

func medConf() appConf {
	constants := []float64{math.E, math.Phi, math.Sqrt(3), math.Sqrt(5), math.Sqrt(7),
		14.1347251417346937904572519835624,
		21.0220396387715549926284795938969,
		25.0108575801456887632137909925628,
		30.4248761258595132103118975305840,
	}

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	lhs := lhsFindsConstants(constants)

	a1 := sequence.Integer{
		Digits:   []int{1, 2, 3, 4, 5},
		Count:    2,
		Repeat:   100,
		Prefix:   []int{1, 2},
		PfxCount: 1,
	}

	a2 := sequence.Polynomial{
		A:    []float64{0, 4},
		B:    []float64{1, 2},
		C:    []float64{0, 2},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 4},
		B:    []float64{-1, 2},
		C:    []float64{0, 2},
		From: 0,
		To:   200,
	}

	cf1 := algorithm.ContinuedFraction{A: a1, B: b}
	cf2 := algorithm.ContinuedFraction{A: a2, B: b}
	nr1 := algorithm.NestedRadical{A: a1, B: b}
	nr2 := algorithm.NestedRadical{A: a2, B: b}

	solvers := []algorithm.Solver{cf1, cf2, nr1, nr2}

	rhs := side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}

	return appConf{
		Constants: constants,
		LHS:       lhs,
		RHS:       rhs,
	}
}

func bigConf() appConf {
	constants := []float64{
		// math.E,
		// math.Phi,
		3.1415926535897932384626433832795028841971693993751058209749445923078, // pi
		// 1.73205080757, // math.Sqrt(3)
		//2.2360679775,  // math.Sqrt(5) comes up a lot
		// 2.64575131106, // math.Sqrt(7)
		14.134725141734693790457251983562470270784257115699243175685567460149, // zeta
		21.022039638771554992628479593896902777334340524902781754629520403587,
		25.010857580145688763213790992562821818659549672557996672496542006745,
		30.424876125859513210311897530584091320181560023715440180962146036993,
		32.935061587739189690662368964074903488812715603517039009280003440784,
		37.586178158825671257217763480705332821405597350830793218333001113622,
		2.92005097731613, // that `prime` constant
		0.00729735256,    // fine structure constant
		137.035999206,
		1.46035450880958681288949915251529801246722933101258149,
	}

	// Find all the constants and their post processed values
	lhs := lhsFindsConstants(constants)

	a1 := sequence.Integer{
		Digits:   []int{1, 2, 3, 4, 5},
		Count:    2,
		Repeat:   100,
		Prefix:   []int{1, 2, 3, 4, 5},
		PfxCount: 1,
	}

	a2 := sequence.Polynomial{
		A:    []float64{-5, 5},
		B:    []float64{-5, 5},
		C:    []float64{-5, 5},
		From: 0,
		To:   200,
	}

	cf1 := algorithm.ContinuedFraction{A: a1, B: a1}
	cf2 := algorithm.ContinuedFraction{A: a1, B: a2}
	cf3 := algorithm.ContinuedFraction{A: a2, B: a1}
	cf4 := algorithm.ContinuedFraction{A: a2, B: a2}
	nr1 := algorithm.NestedRadical{A: a1, B: a1}
	nr2 := algorithm.NestedRadical{A: a2, B: a1}
	nr3 := algorithm.NestedRadical{A: a1, B: a2}
	nr4 := algorithm.NestedRadical{A: a2, B: a2}

	solvers := []algorithm.Solver{cf1, cf2, cf3, cf4, nr1, nr2, nr3, nr4}

	rhs := side{
		Algorithms: solvers,
		PostProc:   true,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}

	return appConf{
		Constants: constants,
		LHS:       lhs,
		RHS:       rhs,
	}
}

func special() appConf {
	constants := []float64{
		2.92005097731613, // that `prime` constant
		0.00729735256,    // another
		137.035999206,
		1.46035450880958681288949915251529801246722933101258149,
	}

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	lhs := lhsFindsConstants(constants)

	a1 := sequence.Integer{
		Digits:   []int{1, 2, 3, 4, 5},
		Count:    2,
		Repeat:   100,
		Prefix:   []int{1, 2, 3, 4, 5},
		PfxCount: 1,
	}

	a2 := sequence.Polynomial{
		A:    []float64{-5, 5},
		B:    []float64{-5, 5},
		C:    []float64{-5, 5},
		From: 0,
		To:   200,
	}

	cf1 := algorithm.ContinuedFraction{A: a1, B: a1}
	cf2 := algorithm.ContinuedFraction{A: a1, B: a2}
	cf3 := algorithm.ContinuedFraction{A: a2, B: a1}
	cf4 := algorithm.ContinuedFraction{A: a2, B: a2}
	nr1 := algorithm.NestedRadical{A: a1, B: a1}
	nr2 := algorithm.NestedRadical{A: a2, B: a1}
	nr3 := algorithm.NestedRadical{A: a1, B: a2}
	nr4 := algorithm.NestedRadical{A: a2, B: a2}

	solvers := []algorithm.Solver{cf1, cf2, cf3, cf4, nr1, nr2, nr3, nr4}

	rhs := side{
		Algorithms: solvers,
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
	}

	return appConf{
		Constants: constants,
		LHS:       lhs,
		RHS:       rhs,
	}
}
