package leetcode

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[0:1]
	}

	return longestString(
		findKindOnePalindrome(s),
		findKindTwoPalindrome(s),
		s[0:1], // if no palindrome has been found, take first letter
	)
}

func findKindOnePalindrome(s string) string {
	var cores []int
	for i := 1; i < len(s)-1; i++ {
		if isKindOnePalindrome(s[i-1 : i+2]) {
			cores = append(cores, i)
		}
	}
	//fmt.Println("K1 cores", cores)

	var result string
	for _, core := range cores {
		p := growKindOnePalindrome(s, core)
		if len(p) > len(result) {
			//fmt.Println("K1", result, p)
			result = p
		}
	}

	return result
}

func isKindOnePalindrome(s string) bool {
	if len(s) != 3 {
		panic("invalid input")
	}
	return s[0] == s[2]
}

func growKindOnePalindrome(s string, middle int) string {
	i := 1
	for {
		// border condition
		if middle-(i+1) < 0 || middle+(i+1) >= len(s) {
			break
		}

		// is a palindrome?
		if s[middle-i-1] != s[middle+i+1] {
			break
		}

		i++
	}
	//fmt.Println(s, middle, i, middle-i, middle+i+1, s[middle-i:middle+i+1])
	return s[middle-i : middle+i+1]
}

func isKindTwoPalindrome(s string) bool {
	if len(s) != 2 {
		panic("invalid input")
	}
	return s[0] == s[1]
}

func findKindTwoPalindrome(s string) string {
	var cores []int
	for i := 0; i < len(s)-1; i++ {
		if isKindTwoPalindrome(s[i : i+2]) {
			cores = append(cores, i)
		}
	}
	//fmt.Println("K2 cores", cores)

	var result string
	for _, core := range cores {
		p := growKindTwoPalindrome(s, core)
		if len(p) > len(result) {
			//fmt.Println("K2", result, p)
			result = p
		}
	}

	return result
}

func growKindTwoPalindrome(s string, start int) string {
	i := 0
	for {
		// border condition
		if start-(i+1) < 0 || start+(i+2) >= len(s) {
			break
		}

		// is a palindrome?
		if s[start-i-1] != s[start+i+2] {
			break
		}

		i++
	}

	return s[start-i : start+i+2]
}

func longestString(ss ...string) string {
	var longest string
	for _, s := range ss {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}
