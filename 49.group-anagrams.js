// @leet start
/**
 * @param {string[]} strs
 * @return {string[][]}
 */
// Time: O(n*m*log*m) (m -> average str length)
// Space: O(n*m)
var groupAnagrams = function(strs) {
    let resMap = {}
	
	for (let str of strs) {
		s = [...str].sort().join('')
		if (resMap[s] === undefined) {
			resMap[s] = [str]
		} else {
			resMap[s].push(str)
		}
	}

	return Object.values(resMap)
};
// @leet end
