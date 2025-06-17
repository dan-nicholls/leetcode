// @leet start
func twoSum(nums []int, target int) []int {
	remainingMap := make(map[int]int, 0)

	for i:= 0; i <= len(nums); i++ {
		remaining := target - nums[i]
		v, ok := remainingMap[remaining]
		if !ok {
			remainingMap[nums[i]] = i
			continue
		}
		return []int{v, i}
	}
	return []int{}
}
// @leet end
