// @leet start
/**
 * @param {number[]} nums
 * @return {number}
 */

// Space Complexity = O(3n) = O(n)
// Time Complexity = O(n*D+n*log(n)*n^2) = O(n^2)
var minSwaps = function(nums) {
	// Create value totals
	let totals = nums.map(x => {
		return x.toString().split("").reduce((acc, x) => acc + parseInt(x),0)
	})
	console.log("totalarr: ", totals)

	// Calculate the correct order
	let idx = nums.map((_, i) => i)

	let sortFunc = (a, b) => {
		let totalA = totals[a]
		let totalB = totals[b]
		if (totalA > totalB) {
			return 1
		}
		if (totalA < totalB) {
			return -1
		}

		return (nums[a] > nums[b]) ? 1 : -1 
	}

	idx.sort(sortFunc)

	// Calc min swaps
	let swaps = 0
	const cur = Array.from({ length: nums.length }, (_,i) => i)

	for (let i=0; i<nums.length; i++) {
		if (cur[i] === idx[i]) { continue }

		let j = i + 1
		while (j < nums.length && cur[j] !== idx[i]) j++;

		[cur[j], cur[i]] = [cur[i],  cur[j]]
		swaps++
	}

	return swaps 
};
// @leet end
