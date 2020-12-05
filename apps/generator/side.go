package main

import (
	"log"
	algo "ramanujan/algorithm"
	"ramanujan/sequence"
)

// SideConf contains all the information to generate a series of calculated values
// from a list of algorithms
type SideConf struct {
	Algorithms []algo.Algorithm
	PostProc   bool
	Ignore     []float64
	ASeqs      []SeqConfig
	BSeqs      []SeqConfig
}

// SeqConfig contains the sequence generator type and its arguments
type SeqConfig struct {
	Generator sequence.Generator
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
func (conf SideConf) Solve() <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)

		for _, algFunc := range conf.Algorithms {
			for _, aConf := range conf.ASeqs {
				for _, bConf := range conf.BSeqs {

					for aseq := range aConf.Generator.Generate() {
						for bseq := range bConf.Generator.Generate() {

							val, err := algFunc.Solve()

							if err != nil {
								log.Fatal("error calling algorithm", err)
							}

							ignore := false
						loop:
							for _, i := range conf.Ignore {
								if val == i {
									ignore = true
									break loop
								}
							}

							if !ignore {
								ch <- val
							}
						}
					}
				}
			}
		}
	}()

	return ch
}
