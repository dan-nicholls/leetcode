// @leet start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(str string) (int, int) {
	zeros, ones := 0, 0
	for _, char := range str {
		if char == '0' {
			zeros++
		} else if char == '1' {
			ones++
		}
	}
	return zeros, ones
}

func findMaxForm(strs []string, m int, n int) int {
	// Initailise 2D array
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Iterate
	for _, str := range strs {
		zeros, ones := count(str)

		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	return dp[m][n]
}

func findMaxForm3D(strs []string, m int, n int) int {
	// Initialise 3D array
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		} 
	}
	
	// Iterate
	for i:=1; i <= len(strs); i++ {
		zeros, ones := count(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; n++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= zeros && k >= ones {
					dp[i][j][k] =  max(dp[i][j][k], dp[i-1][j-zeros][k-ones]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}
// @leet end
