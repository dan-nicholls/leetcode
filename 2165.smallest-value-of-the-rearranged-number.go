// @leet start
func IntToSlice(num int64, numSlice []int64) []int64 {
	if num != 0 {
		i := num % 10
		numSlice = append([]int64{i}, numSlice...)
		return IntToSlice(num/10, numSlice)
	}
	return numSlice
}

func SliceToInt(nums []int64) int64 {
	var result int64 = 0
	for _, i := range nums {
		result = result*10 + i
	}
	return result
}

func SortVals(nums []int64, isAsc bool) []int64 {
	if isAsc {
		slices.Sort(nums)
	} else {
		slices.SortFunc(nums, func (a, b int64) int {
			if a > b {
				return -1
			} else if a < b {
				return 1
			}
			return 0
		})
	}

	if nums[0] == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i] != 0 {
				nums[0], nums[i] = nums[i], nums[0]
				break
			}
		}
	}

	return nums
}



func smallestNumber(num int64) int64 {
	if num >= -1 && num <= 10 {
		return num
	}

	if num < 0 {
		num = num * -1
		
		//Get Largest Number
		numSlice := IntToSlice(num, []int64{}) 
		numSlice = SortVals(numSlice, false)
		num = SliceToInt(numSlice) *-1
		return num
	}

	// Get smallest number
	numSlice := IntToSlice(num, []int64{}) 
	numSlice = SortVals(numSlice, true)
	num = SliceToInt(numSlice)
    
	return num
}
// @leet end
