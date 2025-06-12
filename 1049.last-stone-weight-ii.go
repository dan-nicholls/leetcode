// @leet start
func lastStoneWeightII(stones []int) int {
	// Find DP table length
	var total int
	for _, val := range stones {
		total += val
	}
	target := total / 2

	// Initialise DP table
	dp := make([][]bool, len(stones)+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		if i == 0 {
			dp[i][0] = true
		}
	}

	// Iterate stones for total vals
	for i := 1; i<= len(stones); i++ {
		for j := target; j >= 0; j-- {
			if j >= stones[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-stones[i-1]]
	 		} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Calculate final stone value based on max target
	var maxVal int
	for j := target; j >= 0 ; j-- {
		if dp[len(stones)][j] {
			maxVal = j
			break
		}
	}

	return total - 2*maxVal
}
// @leet end
