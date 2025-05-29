// @leet start
import (
	"fmt"
)

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func findTargetSumWays(nums []int, target int) int {
	// Time Complexity = O(n) + O(n x total) + o(n x total) ≈ O(n x total)
	// Find Max
	max := sum(nums)
	
	// Check target is achievable
	if target > max || target < -max {
		return 0
	}
	
	// Prep DP Table
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, max * 2 + 1)
	}
	dp[0][max] = 1

	// Interate
	for i:=0; i < len(nums); i++ {
		for j := range dp[i] {
			if dp[i][j] != 0 {
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
				fmt.Printf("num: [%d]=%d\n", i, nums[i])
				dp[i+1][j-nums[i]] = dp[i+1][j-nums[i]] + dp[i][j]
				dp[i+1][j+nums[i]] = dp[i+1][j+nums[i]] + dp[i][j]
			}
		}
	}

	// Ensure we return the target + offset
	targetIndex := max + target
	return dp[len(nums)][targetIndex]
}
// @leet end
