// @leet start
func isCovered(ranges [][]int, left int, right int) bool {
	// TODO- how to slice and populate based on range
	isCoveredSlice := make([]bool, right-left+1)

	// Create Slice with values
	for _, currentRange := range ranges {
		start := max(currentRange[0], left)
		end := min(currentRange[1], right)

		for i := start; i <= end; i++ {
			isCoveredSlice[i-left] = true
		}
	}

	// Validate Ranges
	for _, isCovered := range isCoveredSlice {
		if !isCovered {
			return false
		}
	}
	return true
}

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

// @leet end
