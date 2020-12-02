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

// Sequence calculates the sequence of all polynomial results using the coefficiant
// ranges passed in and the values in the range from polyRange[0] to polyRange[1]
// and returns them through a channel
func Sequence(coeffs [][]float64, polyLow, polyHigh float64) <-chan []float64 {
	ch := make(chan []float64)
	coeff := coefficients(coeffs[0], coeffs[1], coeffs[2])

	if polyHigh < polyLow {
		log.Fatal("Invalid argument: polyHigh should be less than polyLow")
	}

	go func() {
		defer close(ch)

		count := int(polyHigh-polyLow) + 1
		for c := range coeff {
			i := 0
			res := make([]float64, count)
			for x := polyLow; x <= polyHigh; x++ {
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
