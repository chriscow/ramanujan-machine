package contfrac

import (
	"log"
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

	res := Solve(a, b)
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

	res := Solve(seq, nil)
	if res != math.Phi {
		t.Log("expected", res, "==", math.Phi)
		t.Fail()
	}
}

func TestCFEuler(t *testing.T) {
	aSeq := make([]float64, 200)
	bSeq := make([]float64, 200)

	a := []float64{3, 1, 0}
	b := []float64{0, -1, 0}

	for i := 0; i < 200; i++ {
		aSeq[i] = poly.Solve(a, float64(i))
		bSeq[i] = poly.Solve(b, float64(i))
	}

	log.Println(aSeq[:5])
	log.Println(bSeq[:5])

	res := Solve(aSeq, bSeq)
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

	log.Println(aSeq[:5])
	log.Println(bSeq[:5])

	res := Solve(aSeq, bSeq)
	if res != math.E {
		t.Log("expected", res, "==", math.E)
		t.Fail()
	}
}
