// @leet start
/**
 * @param {number[]} nums
 * @return {boolean}
 */
// Space Complexity: O(n)
// Time Complexity: O(n)
var containsDuplicate = function(nums) {
	if (nums.length === 1) {
		return false
	}

	let map = {}

	for(let num of nums) {
		console.log(num)
		if (map[num] === undefined) {
			map[num] = true
			continue
		} else if (map[num] === true) {
			return true
		}
	}

	return false
};
// @leet end
