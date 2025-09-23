// @leet start
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var hasTrailingZeros = function(nums) {
	return nums.filter((x) => x % 2 == 0 ? true : false).length >= 2
};
// @leet end
