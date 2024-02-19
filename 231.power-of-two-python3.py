# @leet start
class Solution:
    # Solution 1 - Binary Val Count
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and bin(n)[2:].count("1") == 1

    # Solution 2 - Binary Shift
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and (n & (n - 1)) == 0


# @leet end
