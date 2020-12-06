package main

import (
	algo "ramanujan/algorithm"
	"ramanujan/sequence"
	"runtime"

	"golang.org/x/sync/semaphore"
)

// Side contains all the information to generate a series of calculated values
// from a list of algorithms
type Side struct {
	Algorithms []algo.Solver
	PostProc   bool
	Ignore     []float64
	ASeqs      []sequence.Generator
	BSeqs      []sequence.Generator
}

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(maxWorkers))
)

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
func (conf Side) Solve() <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)

		for _, algFunc := range conf.Algorithms {

			for val := range algFunc.Solve() {
				// see if we are supposed to ignore the returned value
				ignore := false
				for _, i := range conf.Ignore {
					if val == i {
						ignore = true
						break
					}
				}

				if !ignore {
					ch <- val
				}
			}
		}
	}()

	return ch
}
