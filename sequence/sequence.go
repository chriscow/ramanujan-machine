package sequence

// Generator function signature for creating sequences of numbers
type Generator interface {
	Generate() <-chan []float64
}
