// @leet start
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	countMap := make(map[int]int)
	result := []int{}
	

	updateMap := func (nums []int, seen map[int]bool){
		for _,val := range nums {
			if !seen[val] {
				countMap[val]++
				seen[val] = true
			}
		}
	}


	updateMap(nums1, make(map[int]bool))
	updateMap(nums2, make(map[int]bool))
	updateMap(nums3, make(map[int]bool))

	for i, count := range countMap {
		if  count >= 2 {
			result = append(result, i)	
		}	
	}
	return result
}
// @leet end
