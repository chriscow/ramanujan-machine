package main

import (
	"log"
	"math"
	"testing"
)

func solve(conf sideConf) []float64 {
	values := make([]float64, 0)

	for _, alg := range conf.algorithms {
		for _, aConf := range conf.aSeqs {
			for _, bConf := range conf.bSeqs {

				for aseq := range aConf.Generator(aConf.Coeff, aConf.RangeLow, aConf.RangeHigh) {
					for bseq := range bConf.Generator(bConf.Coeff, bConf.RangeLow, bConf.RangeHigh) {

						val, err := alg(aseq, bseq)
						if err != nil {
							log.Fatal("error calling algorithm", err)
						}

					loop:
						for i := range conf.blacklist {
							if val == conf.blacklist[i] {
								break loop
							}
						}

						values = append(values, val)
					}
				}
			}
		}
	}
	return values
}

func TestEandPhiConf(t *testing.T) {
	conf := findsEandPhi()

	e := false
	phi := false

	for _, val := range solve(conf) {
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

func TestEConf(t *testing.T) {
	conf := findsE()

	for _, val := range solve(conf) {
		if val == math.E {
			return
		}
	}

	t.Log("Did not find e")
	t.Fail()
}

func TestPhiCFandNRConf(t *testing.T) {
	conf := findsPhiCFandNR()

	count := 0
	for _, val := range solve(conf) {
		if val == math.Phi {
			count++
		}
	}

	if count != 2 {
		t.Log("expected to find phi twice but found it", count)
		t.Fail()
	}
}
