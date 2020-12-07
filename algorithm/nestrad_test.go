package algorithm

import (
	"math"
	"ramanujan/sequence"
	"testing"
)

func TestNRGetType(t *testing.T) {
	cf := NestedRadical{}
	if cf.GetType() != NestedRad {
		t.Log("expected type NestedRad but got", cf.GetType())
		t.Fail()
	}
}

// TestNRPhi finds the constant phi using a nested radical
func TestNRPhi(t *testing.T) {
	// aSeq := make([]float64, 200)
	// bSeq := make([]float64, 200)

	a := sequence.Polynomial{
		A:    []float64{1, 2},
		B:    []float64{0, 1},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	nr := NestedRadical{A: a, B: a}

	// Solve the nested radical for phi
	for res := range nr.Solve() {
		if res.Result == math.Phi {
			return
		}
	}

	t.Log("not found:", math.Phi)
	t.Fail()
}

func TestRamanNR(t *testing.T) {
	a := sequence.Polynomial{
		A:    []float64{1, 2},
		B:    []float64{0, 1},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	b := sequence.Polynomial{
		A:    []float64{2, 3},
		B:    []float64{1, 2},
		C:    []float64{0, 1},
		From: 0,
		To:   200,
	}

	nr := NestedRadical{A: a, B: b}

	// Finds 3  (Ramanujan’s nested radical)
	for res := range nr.Solve() {
		if res.Result == 3 {
			return
		}
	}

	t.Log("did not find 3")
	t.Fail()
}
