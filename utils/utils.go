package utils

// Reverse reverses the elements of a float64 slice by swapping elements in place
func Reverse(s []float64) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
