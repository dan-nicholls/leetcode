// @leet start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
	return s.length === t.length && [...s].sort().join('') === [...t].sort().join('')
}

// Time: O(n)
// Space: O(k)
var isAnagramMap = function(s, t) {
	if (s.length !== t.length) return false

	let sMap = {}

	for (let c of s) {
		if (sMap[c] == undefined) {
			sMap[c] = 0
		} else {
			sMap[c]++
		}
	}

	for (let c of t) {
		if (sMap[c] === undefined || sMap[c] === 0) {
			return false
		}
		sMap[c]--
	}

	return true
};

// Time: O(n^2) 
// Space: O(n)
var isAnagramLoop = function(s, t) {
	let tCopy = t

	for (let c of s) {
		let cIndex = tCopy.indexOf(c)
		if (cIndex === -1) {
			return false
		}
		tCopy = tCopy.slice(0, cIndex) + tCopy.slice(cIndex+1)
	}

	return tCopy.length === 0
};
// @leet end
