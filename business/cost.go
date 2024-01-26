package business

import "golang.org/x/exp/constraints"

func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (w Wind) Cost() float64 {
	return w.Netto * 0.3
}

// Number is either a floating point number or an integer
type Number interface {
	constraints.Float | constraints.Integer
}

// Cost multiplies usage with netto to compute the cost.
func Cost[T Number](usage, netto T) T {
	cost := usage * netto
	return cost
}
