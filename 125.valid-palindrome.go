// @leet start
// Time = O(n/2) = O(n)
// Space = O(1)
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		// Set valid i
		for i < j && !IsValidAsciiNum(s[i]) {
			i++
		}

		// Set valid j
		for i < j && !IsValidAsciiNum(s[j]) {
			j--
		}

		//Compare
		if asciiLower(s[i]) != asciiLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func IsValidAsciiNum(b byte) bool {
	return ((b >= 'a') && (b <= 'z')) ||
		((b >= 'A') && (b <= 'Z')) ||
		((b >= '0') && (b <= '9'))
}

func asciiLower(b byte) byte {
	if (b >= 'A') && (b <= 'Z') {
		return b + ('a' - 'A') // 32
	}
	return b
}

// Time = O(n+m)
// Space = O(m)
func isPalindromeWithBuilder(s string) bool {
	var finalStr strings.Builder

	for i := 0; i < len(s); i++ {
		b := s[i]
		if (b >= 'A') && (b <= 'Z') {
			b = b + 32
		}
		if ((b >= 'a') && (b <= 'z')) ||
			((b >= '0') && (b <= '9')) {
			finalStr.WriteByte(b)
		}
	}

	s = finalStr.String() // <- reallocation
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

// @leet end
