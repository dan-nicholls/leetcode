// @leet start
func countValidSelections(nums []int) int {
	res := 0
	for _, s := range getInitialMoves(nums) {
		for s.curr >= 0 && s.curr < len(s.nums) {
			step(&s)
		}

		if isValidFinalState(s) {
			res++
		}
	}
	return res
}

type State struct {
	nums      []int
	curr      int
	increment int
}

func step(s *State) {
	if s.nums[s.curr] > 0 {
		s.nums[s.curr]--
		s.increment = -s.increment
	}
	s.curr += s.increment
}

func getInitialMoves(nums []int) []State {
	res := make([]State, 0)
	for i, num := range nums {
		if num == 0 {
			// Ensure that the backing array isnt shared between initial moves
			a := append([]int(nil), nums...)
			b := append([]int(nil), nums...)
			res = append(res,
				State{nums: a, curr: i, increment: 1},
				State{nums: b, curr: i, increment: -1},
			)
		}
	}
	return res
}

func isValidFinalState(s State) bool {
	for _, n := range s.nums {
		if n != 0 {
			return false
		}
	}
	return true
}

// @leet end
