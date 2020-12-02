package rational

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

	res := Solve(a, b)
	if res != math.E {
		t.Log("expected", res, "==", math.E)
		t.Fail()
	}
}
