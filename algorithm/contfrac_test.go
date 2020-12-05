package algorithm

import (
	"math"
	poly "ramanujan/sequence/polynomial"
	"testing"
)

func TestCFPhi(t *testing.T) {
	a := make([]float64, 200)
	b := make([]float64, 200)
	coeff := []float64{1, 0, 0}
	for i := 0; i < 200; i++ {
		a[i] = poly.Solve(coeff, float64(i))
		b[i] = poly.Solve(coeff, float64(i))
	}

	res, err := Solve(a, b)
	if err != nil {
		t.Log("continued fraction phi-1", err)
		t.Fail()
	}

	if res != math.Phi {
		t.Log("expected", res, "==", math.Phi)
		t.Fail()
	}
}

func TestCFPhi2(t *testing.T) {
	seq := make([]float64, 50)
	for i := range seq {
		seq[i] = 1
	}

	res, err := Solve(seq, nil)
	if err != nil {
		t.Log("continued fraction phi-2", err)
		t.Fail()
	}

	if res != math.Phi {
		t.Log("expected", res, "==", math.Phi)
		t.Fail()
	}
}

// TestCFEuler finds the Euler constant using a continued fraction of the form:
//
func TestCFEuler(t *testing.T) {
	aSeq := make([]float64, 200)
	bSeq := make([]float64, 200)

	a := []float64{3, 1, 0}
	b := []float64{0, -1, 0}

	for i := 0; i < 200; i++ {
		aSeq[i] = poly.Solve(a, float64(i))
		bSeq[i] = poly.Solve(b, float64(i))
	}

	res, err := Solve(aSeq, bSeq)
	if err != nil {
		t.Log("continued fraction e #1", err)
		t.Fail()
	}

	if res != math.E {
		t.Log("expected", res, "==", math.E)
		t.Fail()
	}
}

func TestCFEuler2(t *testing.T) {
	aSeq := make([]float64, 147)
	bSeq := make([]float64, 147)

	for i := 3; i < 150; i++ {
		aSeq[i-3] = float64(i)
	}

	for i := 1; i < 148; i++ {
		bSeq[i-1] = float64(-i)
	}

	res, err := Solve(aSeq, bSeq)
	if err != nil {
		t.Log("continued fraction e #2", err)
		t.Fail()
	}

	if res != math.E {
		t.Log("expected", res, "==", math.E)
		t.Fail()
	}
}
