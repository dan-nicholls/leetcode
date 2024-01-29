# @leet start
class Solution:
    def createDpTable(self, target: int, n: int):
        # Create a DP table
        dp = [[0 for _ in range(target + 1)] for _ in range(n + 1)]
        dp[0][0] = 1

        for i in range(1, n + 1):
            for j in range(target + 1):
                for dice_roll in range(1, 7):
                    if j >= dice_roll:
                        dp[i][j] += dp[i - 1][j - dice_roll]
        return dp

    def missingRolls(self, rolls: List[int], mean: int, n: int) -> List[int]:
        target = mean * (len(rolls) + n) - sum(rolls)

        if target < n or target > 6 * n:
            return []

        dp = self.createDpTable(target, n)

        if dp[n][target] == 0:
            return []

        solution = []
        remaining_sum = target
        for i in range(n, 0, -1):
            for roll in range(1, 7):
                if remaining_sum >= roll and dp[i - 1][remaining_sum - roll] > 0:
                    solution.append(roll)
                    remaining_sum -= roll
                    break
        return solution

        # memoCache = {}
        #
        # def findVals(currentSet, numVals, targetTotal):
        #     # Sort key for memoization
        #     currentSetKey = tuple(sorted(currentSet))
        #
        #     if currentSetKey in memoCache:
        #         return memoCache[currentSetKey]
        #
        #     currentTotal = sum(currentSet)
        #
        #     if len(currentSet) == numVals:
        #         result = currentSet if currentTotal == targetTotal else []
        #         memoCache[currentSetKey] = result
        #         return result
        #
        #     if currentTotal > targetTotal or len(currentSet) > numVals:
        #         memoCache[currentSetKey] = []
        #         return []
        #
        #     for number in range(1, 7):
        #         newSet = currentSet + [number]
        #         result = findVals(newSet, numVals, targetTotal)
        #
        #         if result != []:
        #             memoCache[currentSetKey] = result
        #             return result
        #
        #     memoCache[currentSetKey] = []
        #     return []
        #
        # return findVals([], n, totalDif)


# @leet end
