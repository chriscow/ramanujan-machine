package sequence

import "testing"

func TestSolvePolynomial(t *testing.T) {
	// 1 + 2x + 3x^2 where x = 5
	// 1 + 10 + 3 * 25 = 11 + 75 = 86
	coeffs := []float64{1, 2, 3}
	res := polynomial(coeffs, 5)

	if res != 86 {
		t.Log("Expected", res, "== 86")
		t.Fail()
	}
}

func TestSumPolynomialSequence(t *testing.T) {
	a := []float64{-4, 4}
	b := []float64{-3, 3}
	c := []float64{-2, 2}

	var sum float64
	ch := coefficients(a, b, c)
	for coeff := range ch {
		sum += polynomial(coeff, 7)
	}

	if sum != -5472 {
		t.Log("Expected", sum, "== -5472")
		t.Fail()
	}
}
