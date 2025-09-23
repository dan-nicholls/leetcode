// @leet start
/**
 * @param {string} word
 * @return {boolean}
 */

// Space Complexity = O(n)
// Time Complexity = O(n)
var detectCapitalUseSimple = function(word) {
	const allUpper = word === word.toUpperCase()
	const allLower = word === word.toLowerCase()
	const capsFirst = word[0] === word.toUpperCase() && word.slice(1) === word.slice(1).toLowerCase
	return allUpper || allLower || capsFirst
};

// Space Complexity = O(n)
// Time Complexity = O(n)
var detectCapitalUse = function(word) {
	let isUpperCase = (s) => {
		return s.toUpperCase() === s 
	}
	// TRUE if: 
	// 1. All are CAPS
	// 2. all not caps
	// 3. only first letter in word is capital
	
	// if i=0 == caps then
	// 	i=1..n == caps
	// 	or i=1..n != caps
	// else
	// 	if 1..n != caps
	if (word.length === 1) { return true }

	let firstChar = word[0]
	if (isUpperCase(firstChar)) {
		// Get second char
		if (isUpperCase(word[1])) {
			// All must be upper
			return word.split("").slice(1).every((x) => isUpperCase(x))
		}

		// all must be lower
		return word.split("").slice(1).every((x) => !isUpperCase(x))
	}
	//all must be lower
	return word.split("").slice(1).every((x) => !isUpperCase(x))
};
// @leet end
