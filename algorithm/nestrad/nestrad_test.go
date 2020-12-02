package nestrad

import (
	"math"
	"ramanujan/sequence/polynomial"
	"testing"
)

func TestNRPhi(t *testing.T) {
	a := []float64{1, 0, 0}
	b := []float64{1, 0, 0}

	aSeq := make([]float64, 200)
	bSeq := make([]float64, 200)

	for i := range aSeq {
		aSeq[i] = polynomial.Solve(a, float64(i))
		bSeq[i] = polynomial.Solve(b, float64(i))
	}

	// Solve the nested radical for phi
	res, err := Solve(aSeq, bSeq)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if res != math.Phi {
		t.Log("expected", res, "==", math.Phi)
		t.Fail()
	}
}
