// @leet start
func contains(s string, wordDict []string) bool {
	for _, word := range wordDict {
		if word == s {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	// Time complexity = O(len(s) * len(s) * len(wordDict)) = O(n^2xm)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	
	for i:=1; i<=len(s); i++ {
		for j:=0; j < i; j++ {
			if dp[j] && contains(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}

	}

	fmt.Println(dp)
	return dp[len(s)]
}
// @leet end
