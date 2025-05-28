// @leet start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	if a < b && a < c{
		return a
	} else if b < c {
		return b
	}
	return c
}

func contains(val int, nums []int) bool {
	for _, i := range nums {
		if i == val {
			return true
		}
	}
	return false
} 

func mincostTickets(days []int, costs []int) int {
	// Time Complexity = O(max(days[-1])*d)
	dp := make([]int, days[len(days)-1]+1)    
	dp[0] = 0

	for i:=1; i<=len(dp)-1; i++ {
		if contains(i, days) {
			shortPassCost := dp[max(0, i-1)] + costs[0]
			mediumPassCost := dp[max(0, i-7)] + costs[1]
			longPassCost := dp[max(0, i-30)] + costs[2]
			
			dp[i] = min(shortPassCost, mediumPassCost, longPassCost)

		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
// @leet end
