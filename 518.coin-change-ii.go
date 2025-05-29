// @leet start
func change(amount int, coins []int) int {
	// Edge Case
	if amount == 0 {
		return 1
	}

	// Initialise Array
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount + 1)
		dp[i][0] = 1
	}

	// Iterate
	for i:=1; i <= len(coins); i++ {
		for j := 1; j<=amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i-1] {
				dp[i][j] = dp[i][j-coins[i-1]] + dp[i][j]
			}
		}
	}

	return dp[len(coins)][amount]
}
// @leet end
