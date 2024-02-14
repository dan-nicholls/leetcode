# @leet start
from collections import deque


class Solution:
    def rearrangeArray(self, nums: List[int]) -> List[int]:
        posVals = deque()
        negVals = deque()
        result = []

        for val in nums:
            if val > 0:
                posVals.append(val)
            else:
                negVals.append(val)

        for pos, neg in zip(posVals, negVals):
            result += [pos, neg]
        return result


# @leet end
