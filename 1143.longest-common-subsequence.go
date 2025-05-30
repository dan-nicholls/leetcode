// @leet start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	// Initialise DP
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	// Iterate
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1	
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[len(text1)][len(text2)]
}
// @leet end
