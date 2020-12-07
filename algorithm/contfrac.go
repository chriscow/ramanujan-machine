package algorithm

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math"
	"ramanujan/sequence"
	"ramanujan/slice"
)

// ContinuedFraction struct to support serialization of generator config
type ContinuedFraction struct {
	A, B sequence.Generator
}

// Solve calculates the continued fraction using the sequence of elements as
// the numerators and denominators.
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
func (cf ContinuedFraction) Solve() <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)
		for a := range cf.A.Next() {
			for b := range cf.B.Next() {
				res, err := cf.solve(a, b)
				if err != nil {
					log.Fatal("continued fraction misconfigured", err)
				}
				ch <- res
			}
		}
	}()

	return ch
}

func (cf ContinuedFraction) String(a, b []float64) string {

	if len(a) < 5 || len(b) < 5 {
		panic(errors.New("Arguments must both have len() >= 4"))
	}

	res, err := cf.solve(a, b)
	if err != nil {
		panic(err)
	}

	b = b[:4]
	a = a[:4]

	// extract all the signs from the values of b and then make all b positive
	sign := make([]string, 4)
	for i := range b {
		if b[i] > 0 {
			sign[i] = "+"
		} else {
			sign[i] = "-"
			b[i] = math.Abs(b[i])
		}
	}

	type tmpdata struct {
		A      []float64
		B      []float64
		Sign   []string
		Result float64
		Tabs   []string
	}

	tabs := make([]string, 4)
	tabs[0] = "\t"
	tabs[1] = "\t\t"
	tabs[2] = "\t\t\t"
	tabs[3] = "\t\t\t\t"

	dat := tmpdata{A: a, B: b, Sign: sign, Result: res, Tabs: tabs}

	tmpl, err := template.New("tmp").Parse(`{{range $i, $a := .A}}
{{$z := index .A $i}}{{$a}}
	{{end}}
			[...]`)

	// 			{{index .B $i}}
	// {{$a}} {{index .Sign $i}} ------------
	// {{end}}

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, dat); err != nil {
		panic(err)
	}

	return buf.String()
}

func (cf ContinuedFraction) solve(a, b []float64) (float64, error) {
	var res float64
	res = 1

	// if b == nil {
	// 	b = make([]float64, len(a))
	// 	for i := 0; i < len(a); i++ {
	// 		b[i] = 1
	// 	}
	// }

	if len(a) == len(b) {
		if b[0] == 0 {
			b = b[1:]
		}
	}

	if len(a) == len(b)+1 {
		res = a[len(a)-1] // result starts with the end of the `a` sequence
		a = a[:len(a)-1]  // remove one element so they are same length now
	}

	if len(a) != len(b) {
		return math.NaN(), fmt.Errorf("Expected len(a) == len(b) a:%d b:%d", len(a), len(b))
	}

	slice.Reverse(a)
	slice.Reverse(b)

	for i := range a {
		if res == 0 {
			break
		}

		res = a[i] + b[i]/res
	}

	return res, nil
}
