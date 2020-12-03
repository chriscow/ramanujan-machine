package utils

// Reverse reverses the elements of a float64 slice by swapping elements in place
func Reverse(s []interface{}) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func Product(digits []interface{}, count int) <-chan []interface{} {
	ch := make(chan []interface{})
	ix := make([]int, count)

	go func() {
		defer close(ch)

		for {
			res := make([]interface{}, count)

			for i := 0; i < count; i++ {
				res[i] = digits[ix[i]]
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
