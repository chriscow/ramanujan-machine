package main

import (
	"log"
	"ramanujan/algorithm"
	"testing"
	"time"
)

func TestPaulsConstant(t *testing.T) {
	lhs := make(map[string][]algorithm.Solution)
	rhs := make(map[string][]algorithm.Solution)

	conf := paulSpecial()

	start := time.Now()
	process(conf.LHS, lhs)
	process(conf.RHS, rhs)
	elapsed := time.Since(start)
	log.Println("elapsed:", elapsed.Seconds())

	for k := range lhs {
		log.Println("lhs key:", k)
		if rlist, ok := rhs[k]; ok {
			// found in both sides
			log.Println("found", k, "on both sides")
			for i := range rlist {
				log.Println("index:", i, "type:", rlist[i].Type, "result:", rlist[i].Result)
			}
		}
	}
}

// func TestRHSFindsSqrt3EandPhi(t *testing.T) {
// 	constants := []float64{math.E, math.Phi, math.Sqrt(3)}

// 	check := make(map[float64]bool)
// 	for _, val := range constants {
// 		t.Log(val)
// 		check[val] = false
// 	}

// 	side := rhsFindsSqrt3EandPhi()

// 	for val := range side.Solve() {
// 		// t.Log(val)
// 		if _, ok := check[val]; ok {
// 			check[val] = true
// 		}
// 	}

// 	for k, v := range check {
// 		if v != true {
// 			t.Log("did not find", k)
// 			t.Fail()
// 		}
// 	}
// }

// func TestLHSFindsConstants(t *testing.T) {
// 	constants := []float64{math.E, math.Phi}
// 	check := make(map[float64]bool)
// 	for _, val := range constants {
// 		check[val] = false
// 	}

// 	side := lhsFindsConstants(constants)

// 	for val := range side.Solve() {
// 		check[val] = true
// 	}

// 	for k, v := range check {
// 		if v != true {
// 			t.Log("did not find", k)
// 			t.Fail()
// 		}
// 	}
// }

// func TestRHSFindsE(t *testing.T) {
// 	side := rhsFindsE()

// 	for val := range side.Solve() {
// 		if val == math.E {
// 			return
// 		}
// 	}

// 	t.Log("Did not find e")
// 	t.Fail()
// }

// func TestRHSFindsPhiandE(t *testing.T) {
// 	conf := rhsPhiandE()

// 	e := false
// 	phi := false

// 	// get each value computed by solve() and if it equals e or phi, set those
// 	// flags.  If they both get set, we are done
// 	for val := range conf.Solve() {
// 		if val == math.E {
// 			e = true
// 		} else if val == math.Phi {
// 			phi = true
// 		}

// 		if e && phi {
// 			return
// 		}
// 	}

// 	if !phi {
// 		t.Log("Did not find phi")
// 	}

// 	if !e {
// 		t.Log("Did not find e")
// 	}
// 	t.Fail()
// }

// // finds the value phi using both the continued fraction and nested radical
// func TestRHSPhiCFandNR(t *testing.T) {
// 	conf := rhsPhiCFandNR()

// 	count := 0
// 	for val := range conf.Solve() {
// 		if val == math.Phi {
// 			count++
// 		}
// 	}

// 	if count != 2 {
// 		t.Log("expected to find phi twice but found it", count)
// 		t.Fail()
// 	}
// }
