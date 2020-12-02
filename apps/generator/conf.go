package main

import (
	algo "ramanujan/algorithm"
	"ramanujan/algorithm/contfrac"
	"ramanujan/algorithm/nestrad"
	seq "ramanujan/sequence"
	"ramanujan/sequence/polynomial"
)

type sideConf struct {
	algorithms []algo.Algorithm
	postproc   bool
	blacklist  []float64
	aSeqs      []seq.Config
	bSeqs      []seq.Config
}

// Configuration that finds phi for both contined fraction and nested radical
func findsPhiCFandNR() sideConf {
	return sideConf{
		algorithms: []algo.Algorithm{
			contfrac.Solve, // continued fraction
			nestrad.Solve,  // nested radical
		},
		postproc:  false,
		blacklist: []float64{-2, -1, 0, 1, 2},
		aSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{1, 2},
					[]float64{0, 1},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
		bSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{1, 2},
					[]float64{0, 1},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
	}
}

// Configuration that finds e
func findsE() sideConf {
	return sideConf{
		algorithms: []algo.Algorithm{
			contfrac.Solve, // continued fraction
		},
		postproc:  false,
		blacklist: []float64{-2, -1, 0, 1, 2},
		aSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{3, 4},
					[]float64{1, 2},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
		bSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{0, 1},
					[]float64{-1, 0},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
	}
}

// Configuration that generates just enough range to find BOTH phi and e
func findsEandPhi() sideConf {
	return sideConf{
		algorithms: []algo.Algorithm{
			contfrac.Solve,
		},
		postproc:  false,
		blacklist: []float64{-2, -1, 0, 1, 2},
		aSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{1, 4},
					[]float64{0, 2},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
		bSeqs: []seq.Config{
			seq.Config{
				Generator: polynomial.Sequence,
				Coeff: [][]float64{
					[]float64{0, 2},
					[]float64{-1, 1},
					[]float64{0, 1},
				},
				RangeLow:  0,
				RangeHigh: 200,
			},
		},
	}
}
