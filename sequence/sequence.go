package sequence

type Generator func(coeffs [][]float64, rangeLow, rangeHigh float64) <-chan float64

type Config struct {
	Generator Generator
	Coeff     [][]float64 // [[a-range] [b-range] [c-range]]
	RangeLow  float64
	RangeHigh float64
}
