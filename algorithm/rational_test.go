package algorithm

import (
	"math"
	poly "ramanujan/sequence/polynomial"
	"testing"
)

func TestRationalE(t *testing.T) {
	// rational function for e
	coeffA := []float64{0, 1, 0}
	coeffB := []float64{1, 0, 0}
	a := poly.Solve(coeffA, math.E)
	b := poly.Solve(coeffB, math.E)

	res, err := Solve([]float64{a}, []float64{b})
	if err != nil {
		t.Log("rational function e", err)
		t.Fail()
	}

	if res != math.E {
		t.Log("expected", res, "==", math.E)
		t.Fail()
	}
}

func TestWrongArgLen(t *testing.T) {
	_, err := Solve([]float64{1, 2}, []float64{1})
	if err == nil {
		t.Log("expected error with wrong array lengths")
		t.Fail()
	}

	_, err = Solve([]float64{1}, []float64{1, 2})
	if err == nil {
		t.Log("expected error with wrong array lengths")
		t.Fail()
	}

	_, err = Solve([]float64{1, 2}, []float64{1, 2})
	if err == nil {
		t.Log("expected error with wrong array lengths")
		t.Fail()
	}
}

func TestBArgZero(t *testing.T) {
	_, err := Solve([]float64{1}, []float64{0})
	if err == nil {
		t.Log("expected error with b arg == 0")
		t.Fail()
	}
}
