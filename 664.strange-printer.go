// @leet start
func strangePrinterOld(s string) int {
	stack := make([]string,0)
	dp := make([]int, len(s))

	// Initialise starting values
	dp[0] = 1
	stack = append(stack, string(s[0]))
	fmt.Printf("DP: %v, stack: %v\n", dp, stack)

	for i:=1; i<len(dp); i++ {
		inStack := false
		for j,val := range stack {
			if val == string(s[i]) {
				fmt.Println("In stack")
				stack = stack[:j]
				dp[i] = dp[i-1]
				inStack = true
				break
			}
		}
		if !inStack {
			dp[i] = dp[i-1] + 1
			stack = append(stack, string(s[i]))
		}

		fmt.Printf("DP: %v, stack: %v\n", dp, stack)
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func strangePrinter(s string) int {
	n := len(s)

	// Initialise the 2D DP map
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i:= 0; i <= n-length; i++ {
			j := i + length - 1
			dp[i][j] = dp[i][j-1] + 1
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k]+ dp[k+1][j-1])
				}
			}
		}
	}

	return dp[0][n-1]
}
// @leet end
