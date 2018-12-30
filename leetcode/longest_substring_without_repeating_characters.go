package leetcode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var (
		chars = make(map[byte]int)
		start = 0
		max   = 0
	)

	for i := 0; i < len(s); i++ {
		//fmt.Println("i", i, "start", start, "max", max)
		c := s[i]

		// repeating character found
		if j, exists := chars[c]; exists {

			// refresh max value if necessary
			if max < i-start {
				max = i - start
			}

			start = j + 1
			chars[c] = i
			continue
		}

		// put character index to map
		chars[c] = i
	}

	return max
}
