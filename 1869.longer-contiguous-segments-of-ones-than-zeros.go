// @leet start
// Time = O(n)
// Space = O(2)
func checkZeroOnes(s string) bool {
	longest := map[rune]int{
		'1': 0,
		'0': 0,
	}

	count := 0
	curRune := '0'
	for _, r := range s {
		if curRune != r {
			curRune = r
			count = 0
		}

		count++
		if longest[r] < count {
			longest[r] = count
		}
	}
	return longest['1'] > longest['0']
}

// @leet end
