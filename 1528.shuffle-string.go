// @leet start
func restoreString(s string, indices []int) string {
	final := make([]rune, len(indices))

	for i, r := range s {
		final[indices[i]]=r
	}

	return string(final)
}
// @leet end
