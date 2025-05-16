// @leet start
func getLongestSubsequence(words []string, groups []int) []string {
	var pointerA int
	currentStr := []string{}

	if len(words) == 1 {
		return words
	}

	for i, val := range groups {
		if i == 0 {
			fmt.Println("First Digit: ", words[i])
			currentStr = append(currentStr, words[i])
			continue
		} 

		if val != groups[pointerA] {
			currentStr = append(currentStr, words[i])
			pointerA = i
		} 
	}
    return currentStr
}
// @leet end
