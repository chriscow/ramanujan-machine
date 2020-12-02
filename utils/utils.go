package utils

func Equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Reverse reverses the elements of a float64 slice by swapping elements in place
func Reverse(s []float64) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

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
