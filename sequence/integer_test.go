package sequence

import (
	"ramanujan/algorithm"
	"reflect"
	"testing"
)

func TestIntSeq(t *testing.T) {
	comp := make([][]float64, 9)

	comp[0] = []float64{3, 1, 1, 1, 1, 1, 1, 1, 1}
	comp[1] = []float64{3, 1, 2, 1, 2, 1, 2, 1, 2}
	comp[2] = []float64{3, 1, 3, 1, 3, 1, 3, 1, 3}
	comp[3] = []float64{3, 2, 1, 2, 1, 2, 1, 2, 1}
	comp[4] = []float64{3, 2, 2, 2, 2, 2, 2, 2, 2}
	comp[5] = []float64{3, 2, 3, 2, 3, 2, 3, 2, 3}
	comp[6] = []float64{3, 3, 1, 3, 1, 3, 1, 3, 1}
	comp[7] = []float64{3, 3, 2, 3, 2, 3, 2, 3, 2}
	comp[8] = []float64{3, 3, 3, 3, 3, 3, 3, 3, 3}

	i := 0
	seq := Integer{
		Digits:   []int{1, 2, 3},
		Count:    2,
		Repeat:   4,
		Prefix:   []int{3},
		PfxCount: 1,
	}

	for seq := range seq.Generate(args) {
		if !reflect.DeepEqual(seq, comp[i]) {
			t.Log("expected sequences to be the same")
			t.Log(seq, comp[i])
			t.Fail()
		}
		i++
	}
}

func TestCFSeq(t *testing.T) {
	iseq := Integer{
		Digits:   []int{1, 2},
		Count:    2,
		Repeat:   200,
		Prefix:   []int{1},
		PfxCount: 1,
	}

	comp := []float64{3, 1, 2, 1, 2, 1, 2, 1, 2}

	i := 0
	for seq := range iseq.Generate() {
		if i == 1 {
			if !reflect.DeepEqual(seq, comp) {
				t.Log("expected second sequence incorrect")
				t.Log("expected:", comp)
				t.Log("got:", seq)
				t.Fail()
			}

			cf := algorithm.ContinuedFraction{}
			res, err := cf.Solve(seq, nil)
			if err != nil {
				t.Log("continued fraction", err)
				t.Fail()
			}

			if res != 3 {
				t.Log("expected continued fraction == 3 but got", res)
				t.Fail()
			}

			return
		}
	}
}
