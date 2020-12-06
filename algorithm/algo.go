package algorithm

// Solver function signature
type Solver interface {
	Solve() <-chan float64
}
