package main

import (
	"math"
	"testing"
)

func TestLHSFindsConstants(t *testing.T) {
	constants := []float64{math.E, math.Phi}
	check := make(map[float64]bool)
	for _, val := range constants {
		check[val] = false
	}

	side := lhsFindsConstants(constants)

	for val := range side.Solve() {
		check[val] = true
	}

	for k, v := range check {
		if v != true {
			t.Log("did not find", k)
			t.Fail()
		}
	}
}

func TestRHSFindsE(t *testing.T) {
	side := rhsFindsE()

	for val := range side.Solve() {
		if val == math.E {
			return
		}
	}

	t.Log("Did not find e")
	t.Fail()
}

func TestRHSFindsPhiandE(t *testing.T) {
	conf := rhsPhiandE()

	e := false
	phi := false

	// get each value computed by solve() and if it equals e or phi, set those
	// flags.  If they both get set, we are done
	for val := range conf.Solve() {
		if val == math.E {
			e = true
		} else if val == math.Phi {
			phi = true
		}

		if e && phi {
			return
		}
	}

	if !phi {
		t.Log("Did not find phi")
	}

	if !e {
		t.Log("Did not find e")
	}
	t.Fail()
}

// finds the value phi using both the continued fraction and nested radical
func TestRHSPhiCFandNR(t *testing.T) {
	conf := rhsPhiCFandNR()

	count := 0
	for val := range conf.Solve() {
		if val == math.Phi {
			count++
		}
	}

	if count != 2 {
		t.Log("expected to find phi twice but found it", count)
		t.Fail()
	}
}
