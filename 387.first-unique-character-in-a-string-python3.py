from typing import Dict
from collections import Counter


# @leet start
class Solution:
    # 1st Solution - Array
    def firstUniqChar1(self, s: str) -> int:
        seen: Dict[str, int] = {}

        for char in s:
            if char in seen:
                seen[char] += 1
            else:
                seen[char] = 1

        for index, char in enumerate(s):
            if seen[char] == 1:
                return index

        return -1

    # Final Solution - Counter Class
    def firstUniqChar(self, s: str) -> int:
        seen = Counter(s)

        for index in range(len(s)):
            if seen[s[char]] == 1:
                return index

        return -1


# @leet end
