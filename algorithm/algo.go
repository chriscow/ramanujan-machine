package algorithm

type AlgoType int

const (
	// ContFrac represent a contin
	ContFrac AlgoType = iota
	NestedRad
	Rational
)

type Solution struct {
	Type   AlgoType
	Hash   []byte
	Args   []byte
	Result float64
}

// Solver function signature
type Solver interface {
	GetType() AlgoType
	Solve() <-chan (Solution)
}
