package main

import (
	"log"
	"ramanujan/internal/algo"
)

type Algorithm interface {
	Calc(a, b []float64) float64
}

type Generator interface {
	Generate(coeffs [][]float64, rangeLow, rangeHigh float64) <-chan float64
}

type Sequence struct {
	generator Generator
	coeff     [][]float64 // [[a-range] [b-range] [c-range]]
	rangeLow  float64
	rangeHigh float64
}

type SideConf struct {
	algorithms []Algorithm
	postproc   bool
	blacklist  []float64
	aSeqs      []Sequence
	bSeqs      []Sequence
}

func main() {

	rhs := SideConf{
		algorithms: []Algorithm{
			algo.ContinuedFraction{},
		},
		postproc:  false,
		blacklist: []float64{-2, -1, 0, 1, 2},
		aSeqs: []Sequence{
			Sequence{
				generator: algo.PolynomialSequence{},
				coeff: [][]float64{
					[]float64{1, 4},
					[]float64{0, 2},
					[]float64{0, 1},
				},
				rangeLow:  0,
				rangeHigh: 201,
			},
		},
		bSeqs: []Sequence{
			Sequence{
				generator: algo.PolynomialSequence{},
				coeff: [][]float64{
					[]float64{1, 4},
					[]float64{0, 2},
					[]float64{0, 1},
				},
				rangeLow:  0,
				rangeHigh: 201,
			},
		},
	}

	log.Println(rhs)

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
