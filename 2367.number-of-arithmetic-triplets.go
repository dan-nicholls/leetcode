// @leet start
func arithmeticTriplets3Pointers(nums []int, diff int) int {
	count := 0
	numLen := len(nums) - 1

	p1, p2, p3 := 0, 1, 2
	for true {
		p1Val := nums[p1]
		p2Val := nums[p2]
		p3Val := nums[p3]

		if (p3Val-p2Val == diff) && (p2Val-p1Val == diff) {
			count++
		}

		if (p3 == numLen) && (p2 == numLen-1) && (p1 == numLen-2) {
			return count
		}

		if p3 != numLen {
			p3 = p3 + 1
			continue
		}

		if p2 != numLen-1 {
			p2 = p2 + 1
			p3 = p2 + 1
			continue
		}

		if p1 != numLen-2 {
			p1 = p1 + 1
			p2 = p1 + 1
			p3 = p2 + 1
			continue
		}
	}
	return count
}

// Time = O(n*n-1*n-2) = O(n^3)
// Space = O(1)
func arithmeticTriplets3PointersImproved(nums []int, diff int) int {
	count := 0
	numMaxIndex := len(nums) - 1

	for p1 := 0; p1 <= numMaxIndex-2; p1++ {
		for p2 := p1 + 1; p2 <= numMaxIndex-1; p2++ {
			if nums[p2]-nums[p1] != diff {
				continue
			}
			for p3 := p2 + 1; p3 <= numMaxIndex; p3++ {
				if (nums[p3]-nums[p2] == diff) && (nums[p2]-nums[p1] == diff) {
					count++
				}
			}
		}
	}
	return count
}

// Time = O(2n) = O(n)
// Space = O(n)
func arithmeticTriplets(nums []int, diff int) int {
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		seen[num] = true
	}

	count := 0
	for _, num := range nums {
		if seen[num-diff] && seen[num-2*diff] {
			count++
		}
	}

	return count
}

func printArrLocations(a, b, c int, nums []int) {
	label := make([]string, len(nums))
	if a >= 0 && a < len(nums) {
		label[a] = "A"
	}
	if b >= 0 && b < len(nums) {
		label[b] = "B"
	}
	if c >= 0 && c < len(nums) {
		label[c] = "C"
	}

	for i, v := range label {
		if i > 0 {
			fmt.Print("\t")
		}
		if v == "" {
			fmt.Print(" ")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()

	for i, v := range nums {
		if i > 0 {
			fmt.Print("\t")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

// @leet end
