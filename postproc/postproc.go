package postproc

import (
	"log"
	"math"
)

type FuncType int

const (
	Identity FuncType = iota
	Inverse
	Squared
	Cubed
	Quartic
	Quintic
	Sextic
	Heptic
	SquaredInverse
	CubicInverse
	QuarticInverse
	QuinticInverse
	SexticInverse
	HepticInverse
	Sqrt
	SqrtInverse
	Sin
	Cos
	Tan
	Cot
	Exp
	Log
	Log10
	SinInverse
	CosInverse
	TanInverse
	CotInverse
	ExpInverse
	LogInverse
	Log10Inverse
)

// Solution struct supports serialization of result and how we arrived at it
type Solution struct {
	Type   FuncType
	Arg    float64
	Result float64
}

func Enumerate(arg float64) <-chan Solution {
	ch := make(chan Solution)

	go func() {
		defer close(ch)
		for ft := Identity; ft <= Log10Inverse; ft++ {
			res := ft.Solve(arg)
			sln := Solution{
				Type:   ft,
				Arg:    arg,
				Result: res,
			}

			ch <- sln
		}
	}()

	return ch
}

func (t FuncType) Solve(arg float64) float64 {
	switch t {
	case Identity:
		return arg
	case Inverse:
		return inverse(arg)
	case Squared:
		return arg * arg
	case Cubed:
		return math.Pow(arg, 3)
	case Quartic:
		return math.Pow(arg, 4)
	case Quintic:
		return math.Pow(arg, 5)
	case Sextic:
		return math.Pow(arg, 6)
	case Heptic:
		return math.Pow(arg, 7)
	case SquaredInverse:
		return inverse(arg * arg)
	case CubicInverse:
		return inverse(math.Pow(arg, 3))
	case QuarticInverse:
		return inverse(math.Pow(arg, 4))
	case QuinticInverse:
		return inverse(math.Pow(arg, 5))
	case SexticInverse:
		return inverse(math.Pow(arg, 6))
	case HepticInverse:
		return inverse(math.Pow(arg, 7))
	case Sqrt:
		return math.Sqrt(arg)
	case SqrtInverse:
		return inverse(math.Sqrt(arg))
	case Sin:
		return math.Sin(arg)
	case Cos:
		return math.Cos(arg)
	case Tan:
		return math.Tan(arg)
	case Exp:
		return math.Exp(arg)
	case Log:
		return math.Log(arg)
	case Log10:
		return math.Log10(arg)
	case SinInverse:
		return inverse(math.Sin(arg))
	case CosInverse:
		return inverse(math.Cos(arg))
	case TanInverse:
		return inverse(math.Tan(arg))
	case ExpInverse:
		return inverse(math.Exp(arg))
	case LogInverse:
		return inverse(math.Log(arg))
	case Log10Inverse:
		return inverse(math.Log10(arg))
	default:
		log.Fatal("Unimplemented FuncType:", t)
	}

	return math.NaN()
}

func inverse(arg float64) float64 {
	if arg == 0 {
		return math.Inf(1)
	}
	return 1 / arg
}
