package algorithm

// AlgoType is an identifier for an algorithm
type AlgoType int

const (
	// ContinuedFraction algorithm type identifier
	ContinuedFraction AlgoType = iota
	// NestedRadical algorithm type identifier
	NestedRadical
	// Rational function algorithm type identifier
	Rational
)

func (a AlgoType) String() string {
	return [...]string{"continued_fraction", "nested_radical", "rational_func"}[a]
}

// Algorithm function signature
type Algorithm func(a, b []float64) (float64, error)
