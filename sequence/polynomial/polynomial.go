package polynomial

import (
	"log"
	"math"
)

// Solve calculates the polynomial represented by the list of coeffs
// coeffs - polynomial coefficients
// 		 a + b*x + c*x^2 ...
// x      - a value to substitute
func Solve(coeffs []float64, x float64) float64 {
	var sum float64
	for i := range coeffs {
		sum += coeffs[i] * math.Pow(x, float64(i))
	}

	return sum
}

// Args supports (de)serialization of polynomial sequence arguments to an arguments file
type Args struct {
	A     []float64
	B     []float64
	C     []float64
	From  float64
	To    float64
	Range []float64
}

// Sequence calculates the sequence of all polynomial results using the coefficiant
// ranges passed in and the values in the range Args.From -> Args.To
// and returns them through a channel
func Sequence(tmp interface{}) <-chan []float64 {
	args := tmp.(Args)
	ch := make(chan []float64)

	coeff := coefficients(args.A, args.B, args.C)

	if args.To < args.From {
		log.Fatal("Invalid argument: polyHigh should be less than polyLow")
	}

	if args.Range == nil {
		// If specific values were not provided in Range, we generate them
		// going from From to To inclusive incrementing by 1
		count := int(args.To - args.From + 1)
		args.Range = make([]float64, count)
		i := 0
		for val := args.From; val <= args.To; val++ {
			args.Range[i] = val
		}
	}

	go func() {
		defer close(ch)

		for c := range coeff {
			i := 0
			res := make([]float64, len(args.Range))
			for _, x := range args.Range {
				res[i] = Solve(c, x)
				i++
			}

			ch <- res
		}
	}()

	return ch
}

// coefficients returns a channel that returns all possible combinations of
// numbers in the range arguments
func coefficients(aRange, bRange, cRange []float64) <-chan []float64 {
	ch := make(chan []float64)

	go func() {
		defer close(ch)

		for a := aRange[0]; a < aRange[len(aRange)-1]; a++ {
			for b := bRange[0]; b < bRange[len(bRange)-1]; b++ {
				for c := cRange[0]; c < cRange[len(cRange)-1]; c++ {
					ch <- []float64{a, b, c}
				}
			}
		}
	}()

	return ch
}

// func polyrange(aCoeff, bCoeff []float64, low, high float64) <-chan float
