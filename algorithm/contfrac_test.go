package algorithm

import (
	"math"
	"ramanujan/sequence"
	"testing"
)

type mockSeq struct {
	Seq []float64
}

func (m mockSeq) Next() <-chan []float64 {
	ch := make(chan []float64)
	go func() {
		defer close(ch)
		ch <- m.Seq
	}()

	return ch
}

func TestCFPhi(t *testing.T) {
	ones := make([]float64, 200)

	for i := 0; i < 200; i++ {
		ones[i] = 1
	}

	cf := ContinuedFraction{
		A: mockSeq{Seq: ones},
		B: mockSeq{Seq: ones},
	}

	for res := range cf.Solve() {
		if res != math.Phi {
			t.Log("expected", res, "==", math.Phi)
			t.Fail()
		}
	}
}

func TestCFPhi2(t *testing.T) {
	ones := make([]float64, 50)
	for i := range ones {
		ones[i] = 1
	}

	// continued fraction that uses all ones calculates phi
	cf := ContinuedFraction{
		A: mockSeq{Seq: ones},
		B: mockSeq{Seq: ones},
	}

	for res := range cf.Solve() {
		if res != math.Phi {
			t.Log("expected", res, "==", math.Phi)
			t.Fail()
		}
	}
}

// TestCFEuler finds the Euler constant using a continued fraction of the form:
//
func TestCFEuler(t *testing.T) {

	a := sequence.Polynomial{
		A:    []float64{3, 4},
		B:    []float64{0, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{0, 2},
		B:    []float64{-1, 0},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	cf := ContinuedFraction{
		A: a,
		B: b,
	}

	for res := range cf.Solve() {
		if res != math.E {
			t.Log("expected", res, "==", math.E)
			t.Fail()
		}
	}
}

func TestStringPhi(t *testing.T) {
	ones := make([]float64, 50)
	for i := range ones {
		ones[i] = 1
	}

	cf := ContinuedFraction{}
	t.Log(cf.String(ones, ones))
	t.Fail()
}
