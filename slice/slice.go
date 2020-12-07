package slice



func Unique(e []string) []string {
    r := []string{}

    for _, s := range e {
        if !Contains(r[:], s) {
            r = append(r, s)
        }
    }
    return r
}

func Contains(e []string, c string) bool {
    for _, s := range e {
        if s == c {
            return true
        }
    }
    return false
}

// Reverse reverses the elements of a float64 slice by swapping elements in place
func Reverse(s []float64) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}