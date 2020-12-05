package main

import (
	algo "ramanujan/algorithm"
	"ramanujan/sequence"
)

const (
	defaultArgs = "default-args.json"
	argsUsage   = "-args args.json"
)

func checkEnv() error {
	// viper.SetEnvPrefix("raman")
	// viper.AutomaticEnv()

	// TODO: verify required environment variables

	return nil
}

func loadArgs(argsFile string) error {

	// setDefaults()

	// viper.SetConfigFile(argsFile)
	// viper.AddConfigPath(".args")
	// viper.AddConfigPath(".")

	// if err := viper.ReadInConfig(); err != nil {
	// 	// Config file not found; ignore and go with defaults
	// 	log.Println("[generator]", argsFile, "not found. using defaults")
	// 	viper.WriteConfigAs("./" + defaultArgs)
	// }

	return nil
}

//__[ Sample Configurations ]___________________________________________________
//
// The following functions will configure side configurations for interesting
// results.  It will either find some well known constants, good for testing or
// or they may define runs of a certain size for scale / load testing.
//
func tiny() (lhs, rhs SideConf) {

	// This is just enough configuration to find phi and e on the right side
	// and configures the left side for the same
	rhs = rhsFindsE()
	lhs = lhsFindsConstants()

	return
}

func lhsFindsConstants() SideConf {
	return SideConf{
		Algorithms: []algo.Algorithm{algo.RationalFunc{}},
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
		ASeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A: []float64{0, 1}, // args are reversed!
					B: []float64{1, 2}, // 0 + 1x + 0x^2 where x == the const
					C: []float64{0, 1},
				},
			},
		},
		BSeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{ // don't forget the args are reversed
					A: []float64{1, 2},
					B: []float64{0, 1}, // 1 + 0x + 0x^2
					C: []float64{0, 1}, // A +  B + C
				},
			},
		},
	}
}

func rhsPhiandE() SideConf {
	return SideConf{
		Algorithms: []algo.Algorithm{algo.ContinuedFraction{}},
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
		ASeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{1, 4},
					B:    []float64{0, 2},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		},
		BSeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{0, 2},
					B:    []float64{-1, 1},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		}, // BSeqs
	} // SideConf
}

// Configuration that finds phi for both contined fraction and nested radical
func rhsPhiCFandNR() SideConf {
	return SideConf{
		Algorithms: []algo.Algorithm{
			algo.ContinuedFraction{},
			algo.NestedRadical{},
		},
		PostProc: false,
		Ignore:   []float64{-2, -1, 0, 1, 2},
		ASeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{1, 4},
					B:    []float64{0, 2},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		},
		BSeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{0, 2},
					B:    []float64{-1, 1},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		},
	}
}

// Configuration that finds e
func rhsFindsE() SideConf {
	return SideConf{
		Algorithms: []algo.Algorithm{algo.ContinuedFraction{}},
		PostProc:   false,
		Ignore:     []float64{-2, -1, 0, 1, 2},
		ASeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{3, 4},
					B:    []float64{1, 2},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		},
		BSeqs: []SeqConfig{
			SeqConfig{
				Generator: sequence.Polynomial{
					A:    []float64{0, 1},
					B:    []float64{-1, 0},
					C:    []float64{0, 1},
					From: 0,
					To:   200,
				},
			},
		}, // BSeqs
	} // SideConf
}

// Configuration that generates just enough range to find BOTH phi and e
// func findsEandPhi() sideConf {
// 	return sideConf{
// 		algorithms: []algo.Algorithm{
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
