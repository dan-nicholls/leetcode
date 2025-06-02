// @leet start
func countSubarrays(nums []int) int {
	counter := 0

	for i:=0; i < len(nums)-2; i++ {
		valA, valB, valC := nums[i], nums[i+1], nums[i+2]
		if valB % 2 == 0 && valA+valC == valB/2 {
			counter++
		}
	}
	return counter
}
// @leet end
