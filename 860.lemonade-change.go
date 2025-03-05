// @leet start
func lemonadeChange(bills []int) bool {
	billMap := map[int]int{5: 0, 10: 0, 20: 0}

	for _, bill := range bills {
		billMap[bill] = billMap[bill] + 1
	}
	fmt.Println("Start: ", billMap)

	if billMap[10] > billMap[5] {
		fmt.Println("Not enough 5s")
		return false
	}

	// Remove 5 bills
	for range billMap[10] {
		billMap[5] = billMap[5] - 1
	}

	fmt.Println(billMap)
	// Pay 20s with 10s
	for range billMap[20] {
		if billMap[10] > 0 && billMap[5] > 0 {
			billMap[10] = billMap[10] - 1
			billMap[5] = billMap[5] - 1
		} else if billMap[5] > 3 {
			billMap[5] = billMap[5] - 3
		} else {
			return false
		}
	}

	fmt.Println(billMap)
	return true
}

// @leet end
