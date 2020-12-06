package main

import (
	"math"
)

const (
	defaultArgs = "default-args.json"
	argsUsage   = "-args args.json"
)

type appConf struct {
	Constants []float64
	LHS       side
	RHS       side
}

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
// The following create side configurations for interesting
// results.  It will either find some well known constants, good for testing or
// or they may define runs of a certain size for scale / load testing.
//
func tinyConf() appConf {
	constants := []float64{math.E}

	// This is just enough configuration to e on the right side
	// and configures the left side for the same
	rhs := rhsFindsE()
	lhs := lhsFindsConstants(constants)

	return appConf{
		Constants: []float64{math.E},
		LHS:       lhs,
		RHS:       rhs,
	}
}
