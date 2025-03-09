// @leet start
func check(nums []int) bool {
	isShifted := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if isShifted {
				return false
			}

			isShifted = true
		}
	}

	// Check start and end nums are sequential
	if isShifted && (nums[len(nums)-1]-nums[0]) > 1 {
		return false
	}

	return true
}

// @leet end
