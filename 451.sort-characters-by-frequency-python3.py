# @leet start
class Solution:
    def frequencySort(self, s: str) -> str:
        charStore = Counter(s)
        result = ""
        for char, count in charStore.most_common():
            result += char * count
        return result


# @leet end
