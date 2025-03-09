// @leet start
func makeFancyString(s string) string {
	finalString := make([]rune, 0)
	count := 1

	for i := 0; i < len(s); i++ {

		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			finalString = append(finalString, rune(s[i]))
		}
	}

	return string(finalString)
}

// @leet end
