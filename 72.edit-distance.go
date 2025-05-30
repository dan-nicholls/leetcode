// @leet start
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	// Make DP table
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// initialise start vals
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	
	// Initialise start vals
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	// Loop
	for i:=1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	 }
	 return dp[len(word1)][len(word2)]
}
// @leet end
