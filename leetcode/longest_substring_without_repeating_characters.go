package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var (
		symbols = make(map[byte]int)
		start   = 0
		max     = 0
	)
	for i := 0; i < len(s); i++ {
		symbol := s[i]

		// repeating character found
		if j, exists := symbols[symbol]; exists {

			// clear outdated positions from symbols
			for p := start; p <= j; p++ {
				delete(symbols, s[p])
			}

			// refresh max value if necessary
			if max < i-start {
				max = i - start
			}

			start = j + 1
			symbols[symbol] = i

			continue
		}

		// put character index to map
		symbols[symbol] = i
	}

	// case if no repeats have been found
	tail := len(s) - start
	if tail == len(symbols) && tail > max {
		max = tail
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	var (
		n   = len(s)
		ans = 0
		m   = make(map[byte]int)
		j   = 0
		i   = 0
	)
	for ; j < n; j++ {
		symbol := s[j]
		if e, exists := m[symbol]; exists {
			if i < e {
				i = e
			}
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
		m[symbol] = j + 1
	}
	return ans
}
