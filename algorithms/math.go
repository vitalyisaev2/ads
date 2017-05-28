package algorithms

// RaiseNumberToNaturalPower returns a^n for a given a, n integers
// for a O(log N) time.
// [A. Shen, 2011, 1.1.4]
func RaiseNumberToNaturalPower(a, n int) int {
	k := n
	b := 1
	c := a
	for k > 1 {
		if k%2 == 0 {
			k /= 2
			c *= c
		} else {
			k--
			b = b * c
		}
	}
	return b * c
}
