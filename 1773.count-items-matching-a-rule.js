// @leet start
/**
 * @param {string[][]} items
 * @param {string} ruleKey
 * @param {string} ruleValue
 * @return {number}
 */

var countMatchesConcise = function(items, ruleKey, ruleValue) {
	const keyMap = { type: 0, color: 1, name: 2 };
	return items.filter((x) => x[keyMap[ruleKey]] === ruleValue).length
}

// Space Complexity = O(1)
// Time Complexity = O(n)
var countMatches = function(items, ruleKey, ruleValue) {
	const keyMap = {
		"type": 0,
		"color": 1,
		"name": 2
	}

	let count = 0;
	let index = keyMap[ruleKey]

	items.forEach(x => {
		if (x[index] === ruleValue) {
			count++
		}
	})
	return count
};
// @leet end
