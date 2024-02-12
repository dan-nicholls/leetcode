# @leet start
class Solution:
    # Solution 1 - Counter
    def majorityElement1(self, nums: List[int]) -> int:
        return Counter(nums).most_common(1)[0][0]

    # Solution 2 - Sorted List
    def majorityElement2(self, nums: List[int]) -> int:
        return sorted(nums)[len(nums) // 2]

    # Solution 3 - Moores Voting Algorithm
    def majorityElement(self, nums: List[int]) -> int:
        currentCount, currentNum = (0, 0)

        for num in nums:
            if currentCount == 0:
                currentNum = num

            if num == currentNum:
                currentCount += 1
            else:
                currentCount -= 1

        return currentNum


# @leet end
