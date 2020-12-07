package utils


// Product generates the cartesean product from the digits passed in using
// count elements in the output.
//
// Example:
//	- digits: {0,1}
//	- count : 3
//
// Result:
// [0, 0, 0] [0, 0, 1] [0, 1, 0] [0, 1, 1] [1, 0, 0] [1, 0, 1] [1, 1, 0], [1, 1, 1]
func Product(digits []int, count int) <-chan []float64 {
	ch := make(chan []float64)
	ix := make([]int, count)

	go func() {
		defer close(ch)

		for {
			res := make([]float64, count)

			for i := 0; i < count; i++ {
				res[i] = float64(digits[ix[i]])
			}

			ch <- res

			done := true
			for i := len(ix) - 1; i >= 0; i-- {
				ix[i]++
				if ix[i] == len(digits) {
					ix[i] = 0
					continue
				} else {
					done = false
					break
				}
			}

			if done {
				return
			}
		}
	}()

	return ch
}
