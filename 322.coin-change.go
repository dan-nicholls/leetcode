// @leet start
func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i:=0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i < amount + 1; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}


	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
// @leet end
