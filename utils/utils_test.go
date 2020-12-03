package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {
	a := []int{0, 1}
	k := 3
	i := 0

	check := make([][]float64, 8)
	check[0] = []float64{0, 0, 0}
	check[1] = []float64{0, 0, 1}
	check[2] = []float64{0, 1, 0}
	check[3] = []float64{0, 1, 1}
	check[4] = []float64{1, 0, 0}
	check[5] = []float64{1, 0, 1}
	check[6] = []float64{1, 1, 0}
	check[7] = []float64{1, 1, 1}

	for res := range Product(a, k) {
		fmt.Println(res)
		for j := range res {
			if res[j] != check[i][j] {
				t.Log("expected res[", j, "] == check[", i, "][", j, "]", res[j], check[i][j])
				t.Fail()
			}
		}
		i++
	}
}

func TestReverse(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	comp := []float64{4, 3, 2, 1}

	Reverse(a)

	if !reflect.DeepEqual(a, comp) {
		t.Log("expected a:", a, "to get reversed to ", comp)
		t.Fail()
	}
}
