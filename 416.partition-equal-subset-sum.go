// @leet start
func sum(nums []int) int {
	var sum int
	for _, val := range nums {
		sum = sum + val	
	}
	return sum
}

func canPartition(nums []int) bool {
	//Prep total
	total := sum(nums)
	if total % 2 != 0 {
		return false
	}
	goalSum := total / 2
	
	//Prep table
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, goalSum + 1)
		dp[i][0] = true
	}

	//Start iterating
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= goalSum; j++ {
			// Check if j can be created using subarray
			dp[i][j] = dp[i-1][j]

			if nums[i-1] == j {
				dp[i][j] = true
			}
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][goalSum]
}
// @leet end
