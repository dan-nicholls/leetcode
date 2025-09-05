// @leet start
/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
	let res = s.toLowerCase().replaceAll(/[^a-z0-9]/gi, "")
	
	if (res.length == 1) {
		return true
	}

	for (let i = 0; i <=  res.length/2; i++) {
		if (res[i] !== res[res.length-1-i]) {
			return false
		}
	}

	return true 
};
// @leet end
