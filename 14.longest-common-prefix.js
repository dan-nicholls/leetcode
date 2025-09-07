// @leet start
/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
	// Catch edge case
	if (strs.length == 1) {
		return strs[0]
	}

	let resultStr = ""
	for (let i in strs[0]) {
		for (let str of strs) {
			console.log(`i:${i}(${strs[0][i]})\t${str}`)
			if (str[i] != strs[0][i]) {
				return resultStr
			}
		}
		resultStr = resultStr + strs[0][i]
	}

	return ""
};
// @leet end
