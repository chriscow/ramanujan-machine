package slice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	comp := []float64{4, 3, 2, 1}

	Reverse(a)

	if !reflect.DeepEqual(a, comp) {
		t.Log("expected a:", a, "to get reversed to ", comp)
		t.Fail()
	}
}
