# @leet start
class Solution:
    def frequencySort(self, s: str) -> str:
        charStore = Counter(s)
        return [char for char, _ in charStore.most_common()]


# @leet end
