# @leet start
class Solution:
    def missingRolls(self, rolls: List[int], mean: int, n: int) -> List[int]:
        totalDif = mean * (len(rolls) + n) - sum(rolls)

        memoCache = {}

        def findVals(currentSet, numVals, targetTotal):
            # Sort key for memoization
            currentSetKey = tuple(sorted(currentSet))

            if currentSetKey in memoCache:
                return memoCache[currentSetKey]

            currentTotal = sum(currentSet)

            if len(currentSet) == numVals:
                result = currentSet if currentTotal == targetTotal else []
                memoCache[currentSetKey] = result
                return result

            if currentTotal > targetTotal or len(currentSet) > numVals:
                memoCache[currentSetKey] = []
                return []

            for number in range(1, 7):
                newSet = currentSet + [number]
                result = findVals(newSet, numVals, targetTotal)

                if result != []:
                    memoCache[currentSetKey] = result
                    return result

            memoCache[currentSetKey] = []
            return []

        return findVals([], n, totalDif)


# @leet end
