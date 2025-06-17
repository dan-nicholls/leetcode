// @leet start
func areAlmostEqual(s1 string, s2 string) bool {
	incorrectVals := make([]int,0)
	for i := range s1 {
		if s1[i] != s2[i] {
			if len(incorrectVals) >= 2 {
				return false	
			}
			incorrectVals = append(incorrectVals, i)
		}
	}

	if len(incorrectVals) == 0 || len(incorrectVals) == 2 && (s1[incorrectVals[0]] == s2[incorrectVals[1]]) && (s2[incorrectVals[0]] == s1[incorrectVals[1]]){
		return true	
	}
	return false
}
// @leet end
