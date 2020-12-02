package utils

import "fmt"

// Reverse reverses the elements of a float64 slice by swapping elements in place
func Reverse(s []float64) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func Product(digits []int, count int) <-chan []int {
	ch := make(chan []int)
	ix := make([]int, count)

	go func() {
		defer close(ch)

		pos := count - 1
		for {
			res := make([]int, count)

			if pos < 0 {
				return
			}

			for i := 0; i < count; i++ {
				res[i] = digits[ix[i]]
			}

			fmt.Println("sending", res, ix)
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
