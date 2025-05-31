// @leet start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		// not enough chars
		return false
	}
	// Initialise DP table
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// Set base cases
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		if s1[:i] == s3[:i] {
			dp[i][0] = true
		}
	}

	for j:= 1; j < len(dp[0]); j++ {
		if s2[:j] == s3[:j] {
			dp[0][j] = true
		}
	}

	// Iterate
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
// @leet end
