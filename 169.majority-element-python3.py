# @leet start
class Solution:
    # Solution 1 - Counter
    def majorityElement1(self, nums: List[int]) -> int:
        return Counter(nums).most_common(1)[0][0]

    # Solution 2 - Sorted List
    def majorityElement(self, nums: List[int]) -> int:
        return nums.sort()[len(nums) // 2]


# @leet end
