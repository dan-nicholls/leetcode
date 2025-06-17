// @leet start
func containsDuplicate(nums []int) bool {
	countMap := make(map[int]bool)
	for i := range nums {
		fmt.Println(i)
		_,ok := countMap[nums[i]]
		if !ok {
			countMap[nums[i]] = true
			continue
		}
		return true
	}
	return false
}
// @leet end
