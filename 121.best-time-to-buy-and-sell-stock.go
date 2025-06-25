// @leet start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buyPrice := prices[0] 
	maxProfit := 0


	for _, p := range prices {
		if p < buyPrice {
			buyPrice = p
		} else {
			curProf := p - buyPrice
			if curProf > maxProfit {
				maxProfit = curProf
			}
		}
	}
	return maxProfit
}
// @leet end
