# @leet start
from collections import deque


class Solution:
    # Solution 1 - Deques
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

    # Solution 2 - Two Pointers
    def rearrangeArray(self, nums: List[int]) -> List[int]:
        result = [0] * len(nums)
        posPointer = 0
        negPointer = 1

        for val in nums:
            if val > 0:
                result[posPointer] = val
                posPointer += 2
            else:
                result[negPointer] = val
                negPointer += 2

        return result


# @leet end
