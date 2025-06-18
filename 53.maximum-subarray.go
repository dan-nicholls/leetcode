// @leet start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	currentMax := nums[0]
	totalMax := nums[0]
	for i := 1; i<len(nums); i++ {
		currentMax = max(nums[i], currentMax+nums[i])
		totalMax = max(totalMax, currentMax)
	}
	return totalMax
}
// @leet end
