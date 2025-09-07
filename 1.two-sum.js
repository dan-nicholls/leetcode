// @leet start
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
	let map = {}
	for (let i in nums) {
		num = nums[i]
		console.log(`i:${i} (${num})`)

		remaining = target - num
		if (map[remaining] == undefined) {
			map[num] = i
			continue
		}
		return [parseInt(map[remaining]), parseInt(i)]
	}
	return [0, 0]
};
// @leet end
