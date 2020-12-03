package algorithm

import (
	"fmt"
	"ramanujan/algorithm/contfrac"
	"ramanujan/algorithm/nestrad"
	"ramanujan/algorithm/rational"
)

// AlgoType is an identifier for an algorithm
type AlgoType int

const (
	// ContinuedFraction algorithm type identifier
	ContinuedFraction AlgoType = iota
	// NestedRadical algorithm type identifier
	NestedRadical
	// RationalFunc algorithm type identifier
	RationalFunc
)

func (a AlgoType) String() string {
	return [...]string{"continued_fraction", "nested_radical", "rational_func"}[a]
}

// Algorithm function signature
type Algorithm func(a, b []float64) (float64, error)

// Get returns an algorithm function from an AlgoType
func Get(t AlgoType) (Algorithm, error) {
	switch t {
	case ContinuedFraction:
		return contfrac.Solve, nil
	case NestedRadical:
		return nestrad.Solve, nil
	case RationalFunc:
		return rational.Solve, nil
	default:
		return nil, fmt.Errorf("unknown algorithm type: %d", t)
	}
}
