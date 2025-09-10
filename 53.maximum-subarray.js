// @leet start
/**
 * @param {number[]} nums
 * @return {number}
 */
// Space Complexity = O(1)
// Time Complexity = O(n)
var maxSubArray = function(nums) {
	let maxTotal = nums[0];
	let currentTotal = 0;

	for(let num of nums) {
		currentTotal = Math.max(num, currentTotal + num)
		maxTotal = Math.max(currentTotal, maxTotal)
	}
	return maxTotal
};
// @leet end
