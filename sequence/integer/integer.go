package integer

import "ramanujan/utils"

// Sequence generates all possible integer sequences
//
// Arguments:
// 	digits - The digits to use to generate the primary part of the sequence
// 	count - how many digits to use from 'digits' that repeat in the sequence
// 	repeat  - how many times to repeat the above sequence
// 	prefix - Just like above but will prefix the sequence with another sequence
// 	prrepeat - how many digits to use from 'prefix_digits' to create the prefix
//
// Usage:
// 	To generate the sequence (among similar others):
// 		3, 1, 2, 1, 2, 1, 2, 1, 2
//
//	digits := []float64{1,2,3}
//	prefix := []float64{3}
// 	for seq := range integer.Sequence([1,2,3], 2, 4, [3], 1) {}
//
//	returns:
//      count
//      ____
//			  repeat
//			  ____, ____, ____
// 	(3, 1, 1, 1, 1, 1, 1, 1, 1)
// 	(3, 1, 2, 1, 2, 1, 2, 1, 2) << ---
// 	(3, 2, 1, 2, 1, 2, 1, 2, 1)
// 	(3, 2, 2, 2, 2, 2, 2, 2, 2)
func Sequence(digits []int, count, repeat int, prefix []int, prcount int) <-chan []float64 {
	ch := make(chan []float64)

	go func() {
		defer close(ch)

		for pfx := range utils.Product(prefix, prcount) {
			for pattern := range utils.Product(digits, count) {
				// need a new slice each time since they are reference vals
				res := make([]float64, 0, len(pfx)+count*repeat)
				res = append(res, pfx...)
				for i := 0; i < repeat; i++ {
					res = append(res, pattern...)
				}

				ch <- res
			}
		}
	}()

	return ch
}
