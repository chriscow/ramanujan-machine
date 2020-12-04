package sequence

import (
	"errors"
	"fmt"
	"ramanujan/sequence/integer"
	"ramanujan/sequence/polynomial"
	"strings"
)

// GeneratorType function type identifier
type GeneratorType int

const (
	// Polynomial sequence generator
	Polynomial GeneratorType = iota
	// Integer sequence generator
	Integer
)

// GeneratorTypes is a convenience array of all GeneratorType values
var GeneratorTypes = []GeneratorType{Polynomial, Integer}

func (g GeneratorType) String() string {
	return [...]string{"polynomial", "integer"}[g]
}

// From returns a GeneratorType from a string representation
func From(str string) (GeneratorType, error) {
	switch strings.ToLower(str) {
	case Polynomial.String():
		return Polynomial, nil
	case Integer.String():
		return Integer, nil
	default:
		return 0, errors.New(fmt.Sprint("unknown type:", str))
	}
}

// Get returns a concrete generator function from a GeneratorType
func Get(t GeneratorType) (Generator, error) {
	switch t {
	case Polynomial:
		return polynomial.Sequence, nil
	case Integer:
		return integer.Sequence, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", t.String())
	}
}

// Generator function signature for creating sequences of numbers
type Generator func(args interface{}) <-chan []float64
