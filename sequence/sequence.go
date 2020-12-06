package sequence

// Generator function signature for creating sequences of numbers
type Generator interface {
	Next() <-chan []float64
}
