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

func GetGenerator(t GeneratorType) (Generator, error) {
	switch t {
	case Polynomial:
		return polynomial.Sequence, nil
	case Integer:
		return integer.Sequence, nil
	default:
		return nil, errors.New(fmt.Sprint("unknown type:", t.String()))
	}
}

type GeneratorArgs interface{}

// Generator function signature for creating sequences of numbers
type Generator func(args GeneratorArgs) <-chan []float64

// Config structure for generators
type Config struct {
	Generator Generator
	Coeff     [][]float64 // [[a-range] [b-range] [c-range]]
	RangeLow  float64
	RangeHigh float64
}
