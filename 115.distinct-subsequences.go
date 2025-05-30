// @leet start
func numDistinct(s string, t string) int {
	dp := make([][]int32, len(s)+1)
	for i := range dp {
		dp[i] = make([]int32, len(t)+1)
	}

	// initialise dp[0]
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}
	return int(dp[len(s)][len(t)])
}
// @leet end
