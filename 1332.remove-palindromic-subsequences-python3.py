# @leet start
class Solution:
    def removePalindromeSub(self, s: str) -> int:
        # Since you can remove a subsequence and not a substring, if the result is not already a palendrome you can remove the subsequence comprising of a single 'a' or 'b' character. This results in the operation taking a maximum of 2 moves.
        if len(s) == 0:
            return 0

        return 1 if s == s[::-1] else 2


# @leet end
