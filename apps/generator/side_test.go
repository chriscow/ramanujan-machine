package main

import (
	"math"
	"testing"
)

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
