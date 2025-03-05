// @leet start

func distinctIntegers(n int) int {
	// n % (n-1) == 1, all numbers get added
	if n == 1 {
		return 1
	}

	return n - 1
}

// @leet end
