package main

// @leet start

func isPalindrome(s string) bool {
	// Use for Palendrome Logic
	pointerA, pointerB := 0, len(s)-1
	
	for pointerA < pointerB {
		if s[pointerA] != s[pointerB] {
			return false
		}
		pointerA++
		pointerB--
	}
	return true
}


func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	for i:=len(s); i>0; i-- {
		start := 0
		for start + i <= len(s) {
			strSlice := s[start: start+i] 
			if isPalindrome(strSlice) {
				return strSlice
			}
			start++
		}
	}
    return ""
}
// @leet end
