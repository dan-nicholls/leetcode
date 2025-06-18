// @leet start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {
	// Get Max Sum
	total := 0
	currentMax := 0
	currentMin := 0
	maxSum := nums[0]
	minSum := nums[0]
	
	for _, val := range nums {
		total += val

		currentMax = max(val, currentMax+val)
		maxSum = max(maxSum, currentMax)

		currentMin = min(val, currentMin+val)
		minSum = min(minSum, currentMin)
	}

	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, total-minSum)
	
}
// @leet end
