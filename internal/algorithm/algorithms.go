package algo

import "log"

type ContinuedFraction struct{}

// Calc calculates the continued fraction
//
// Parameters:
// a - a list of additive constant (first element) + denominator values
// 	(following elements)
//
// b - a list of numerator values. If ommited, 1's are used.
//
// If a and b are given and are of the same length, the additive constant is
// assumed to be 0 and is all denominators and the result is
//
// b[0] / (a[0] + (b[1] / (a[1] + (b[2] / a[2] + ...
//
// otherwise the result is
//
// a[0] + b[0] / (a[1] + b[1] / (a[2] + b[2] / a[3] ...))
func (cf ContinuedFraction) Calc(a, b []float64) float64 {
	var res float64
	res = 1

	if b == nil {
		b = make([]float64, len(a))
		for i := 0; i < len(a); i++ {
			b[i] = 1
		}
	} else if len(a) == len(b) {
		if b[0] == 0 {
			b = b[1:]
		}
	}

	if len(a) == len(b)+1 {
		res = a[len(a)-1] // result starts with the end of the `a` sequence
		a = a[:len(a)-1]  // remove one element so they are same length now
	}

	if len(a) != len(b) {
		log.Fatal("Expected len(a) == len(b) a:", len(a), "b:", len(b))
	}

	reverse(a)
	reverse(b)

	for i := range a {
		if res == 0 {
			break
		}

		res = a[i] + b[i]/res
	}

	return res
}

func reverse(s []float64) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
