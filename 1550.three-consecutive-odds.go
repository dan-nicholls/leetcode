// @leet start
func threeConsecutiveOdds(arr []int) bool {
	count := 0

	for i := 0; i < len(arr); i++ {
		fmt.Printf("[%d] = %d", i, arr[i])
			
		if (arr[i] % 2 != 0) {
			count++
			fmt.Printf("ODD %d\n", count)
			if count >=3 {
				return true
			}
		} else {
			count = 0
			fmt.Printf("EVEN %d\n", count)
		}
	}
	return false
}
// @leet end
