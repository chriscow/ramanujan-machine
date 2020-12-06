package sequence

// Constant is expected to hold a list of "interesting" constants. I know it
// appears elaborate just to return a single array of values but I want it to
// match the signature of other generator functions
type Constant struct {
	Values []float64
}

// Next returns the next constant value on each iteration
func (c Constant) Next() <-chan []float64 {
	ch := make(chan []float64)

	go func() {
		defer close(ch)
		ch <- c.Values
	}()

	return ch
}
