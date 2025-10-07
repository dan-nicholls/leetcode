// @leet start
// Space Compexity = O(n)
// Time Complexity = O(n)
func plusOne(digits []int) []int {
	for i := len(digits)-1; i>=0; i-- {
		//fmt.Printf("%d - %d\n", i, digits[i] )
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}

	//final := make([]int, len(digits) +1)
	//final[0] = 1
	digits = append([]int{1}, digits...)
	return digits 
}
// @leet end
