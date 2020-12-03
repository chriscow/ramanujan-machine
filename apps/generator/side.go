package main

import (
	"log"
	algo "ramanujan/algorithm"
	"ramanujan/sequence"
)

// SideConf contains all the information to generate a series of calculated values
// from a list of algorithms
type SideConf struct {
	Algorithms []algo.AlgoType
	PostProc   bool
	Ignore     []float64
	ASeqs      []SeqConfig
	BSeqs      []SeqConfig
}

// SeqConfig contains the sequence generator type and its arguments
type SeqConfig struct {
	Generator sequence.GeneratorType
	Args      interface{}
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

		for _, algType := range conf.Algorithms {
			for _, aConf := range conf.ASeqs {
				for _, bConf := range conf.BSeqs {

					aGenFunc, err := sequence.Get(aConf.Generator)
					if err != nil {
						log.Fatal("get a-sequence generator:", err)
					}

					bGenFunc, err := sequence.Get(bConf.Generator)
					if err != nil {
						log.Fatal("get b-sequence generator:", err)
					}

					algFunc, err := algo.Get(algType)
					if err != nil {
						log.Fatal("get algorithm:", err)
					}

					for aseq := range aGenFunc(aConf.Args) {
						for bseq := range bGenFunc(bConf.Args) {
							a := aseq
							b := bseq

							val, err := algFunc(a, b)
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
