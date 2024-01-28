# @leet start
class Solution:
    def missingRolls(self, rolls: List[int], mean: int, n: int) -> List[int]:
        # TODO - Add support for memoization
        totalDif = mean * (len(rolls) + n) - sum(rolls)

        def findVals(currentSet, numVals, targetTotal):
            currentTotal = sum(currentSet)

            if len(currentSet) == numVals:
                return currentSet if currentTotal == targetTotal else []

            if currentTotal > targetTotal or len(currentSet) > numVals:
                return []

            for number in range(1, 7):
                newSet = currentSet + [number]
                result = findVals(newSet, numVals, targetTotal)

                if result != []:
                    return result

            return []

        return findVals([], n, totalDif)


# @leet end
