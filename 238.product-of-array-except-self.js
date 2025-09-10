// @leet start
/**
 * @param {number[]} nums
 * @return {number[]}
 */
// Space Complexity = O(1) (Just the output arr)
// Time Complexity = O(2n) = O(n)
var productExceptSelf = function(nums) {
	const result = new Array(nums.length)
	result[0] = 1;

	for(let i=1;i<nums.length; i++) {
		console.log(`nums[${i}] = ${nums[i]}`)
		result[i] = result[i-1] * nums[i-1]
	}

	currentVal = 
	for(let i=nums.length-1; i>=0; i--) {
		console.log(`nums[${i}] = ${nums[i]}`)
		result[i] = result[i] * currentVal
		currentVal = currentVal * nums[i]
	}

	return result
};
// @leet end
