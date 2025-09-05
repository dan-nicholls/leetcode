// @leet start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
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
