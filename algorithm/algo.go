package algorithm

// Algorithm function signature
type Algorithm interface {
	Solve() (float64, error)
}
