package main

import (
	"log"
	algo "ramanujan/algorithm"
	"ramanujan/algorithm/contfrac"
	"ramanujan/algorithm/nestrad"
	seq "ramanujan/sequence"
	"ramanujan/sequence/polynomial"

	"github.com/spf13/viper"
)

type sideConf struct {
	algorithms []algo.Algorithm
	postproc   bool
	blacklist  []float64
	aSeqs      []seq.Config
	bSeqs      []seq.Config
}

const (
	defaultArgs = "default-args.json"
	argsUsage   = "-args args.json"
)

func checkEnv() error {
	viper.SetEnvPrefix("raman")
	viper.AutomaticEnv()

	// TODO: verify required environment variables

	return nil
}

func loadArgs(argsFile string) error {

	setDefaults()

	viper.SetConfigFile(argsFile)
	viper.AddConfigPath(".args")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// Config file not found; ignore and go with defaults
		log.Println("[generator]", argsFile, "not found. using defaults")
		viper.WriteConfigAs("./" + defaultArgs)
	}

	return nil
}

func setDefaults() {
	//
	// Set default values if no argument file is passed.  Also used to
	// create a basic argument file to start from.  These defaults generate
	// a range to find both phi and e
	//
	// use string values for constants so we can get them at varying decimal
	// precision at runtime, or read them from files
	viper.SetDefault("constants", []string{"e", "phi"})

	//__[ Left Hand Side ]______________________________________________________
	//
	// algorithms to calculate left hand side equation values
	viper.SetDefault("lhs.algorithms", []algo.AlgoType{algo.Rational})
	// run post processing calculations on resulting values from above
	viper.SetDefault("lhs.run_postproc", false)
	// ignore results that equal these values (you get a lot of them)
	viper.SetDefault("lhs.black_list", []float64{-2, -1, 0, 1, 2})
	// algorithms arguments that are generally sequences of numbers.
	// algorithms take two sequences, a and b here.
	//
	// a-sequences
	viper.SetDefault("lhs.a_sequences.generator", seq.Polynomial)
	viper.SetDefault("lhs.a_sequences.arguments", [][]float64{
		[]float64{0, 1},
		[]float64{1, 2},
		[]float64{0, 1},
		nil,
	})
	//
	// b-sequences
	viper.SetDefault("lhs.b_sequences.generator", seq.Polynomial)
	viper.SetDefault("lhs.b_sequences.arguments", [][]float64{
		[]float64{1, 2},
		[]float64{0, 1},
		[]float64{0, 1},
		nil,
	})

	//__[ Right Hand Side ]______________________________________________________
	//
	// algorithms to calculate right hand side equation values
	viper.SetDefault("rhs.algorithms", []algo.AlgoType{algo.ContinuedFraction})
	// run post processing calculations on resulting values from above
	viper.SetDefault("rhs.run_postproc", false)
	// ignore results that equal these values (you get a lot of them)
	viper.SetDefault("rhs.black_list", []float64{-2, -1, 0, 1, 2})
	// algorithms arguments that are generally sequences of numbers.
	// algorithms take two sequences, a and b here.
	//
	// a-sequences
	viper.SetDefault("rhs.a_sequences.generator", seq.Polynomial)
	viper.SetDefault("rhs.a_sequences.arguments", [][]float64{
		[]float64{1, 4},
		[]float64{0, 2},
		[]float64{0, 1},
		[]float64{0, 200},
	})
	//
	// b-sequences
	viper.SetDefault("rhs.b_sequences.generator", seq.Polynomial)
	viper.SetDefault("rhs.b_sequences.arguments", [][]float64{
		[]float64{0, 2},
		[]float64{-1, 1},
		[]float64{0, 1},
		[]float64{0, 200},
	})
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
