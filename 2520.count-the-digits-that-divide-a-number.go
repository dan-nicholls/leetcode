// @leet start
package main
import (
	"strconv"
)
func countDigits(num int) int {
	// Can improve by applying a map for memoization
	nums := strconv.Itoa(num)
	count := 0
	for _,val := range nums {
		if num % int(val -'0') == 0 {
			count++
		}
	}
    return count
}
// @leet end
