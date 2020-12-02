package main

import (
	algo "ramanujan/algorithm"
	seq "ramanujan/sequence"
)

type SideConf struct {
	algorithms []algo.Algorithm
	postproc   bool
	blacklist  []float64
	aSeqs      []seq.Config
	bSeqs      []seq.Config
}

func main() {

	// rhs := SideConf{
	// 	algorithms: []algo.Algorithm{
	// 		algo.ContinuedFraction,
	// 	},
	// 	postproc:  false,
	// 	blacklist: []float64{-2, -1, 0, 1, 2},
	// 	aSeqs: []seq.Config{
	// 		seq.Config{
	// 			Generator: seq.Polynomial,
	// 			Coeff: [][]float64{
	// 				[]float64{1, 4},
	// 				[]float64{0, 2},
	// 				[]float64{0, 1},
	// 			},
	// 			RangeLow:  0,
	// 			RangeHigh: 201,
	// 		},
	// 	},
	// 	bSeqs: []seq.Config{
	// 		seq.Config{
	// 			Generator: seq.Polynomial,
	// 			Coeff: [][]float64{
	// 				[]float64{1, 4},
	// 				[]float64{0, 2},
	// 				[]float64{0, 1},
	// 			},
	// 			RangeLow:  0,
	// 			RangeHigh: 201,
	// 		},
	// 	},
	// }

	// count := 0
	// for alg := range rhs.algorithms {
	// 	count++
	// 	for _, aSeq := range rhs.aSeqs {
	// 		for _, bSeq := range rhs.bSeqs {

	// 			aGen := aSeq.generator
	// 			bGen := bSeq.generator

	// 		}
	// 	}
	// }

}
