# @leet start
class Solution:
    def minimumRecolors(self, blocks: str, k: int) -> int:
        startPtr = 0
        endPrt = k

        bestCount: Tuple(int, int) = None

        print("Blocks: ", blocks)
        print("k: ", k)

        for i in range(len(blocks) - k + 1):
            print("range: ", blocks[i : i + k])
            windowStr = blocks[i : i + k]
            count = windowStr.count("B")

            if (bestCount is None) or (bestCount[1] < count):
                bestCount = (i, bestCount)
            print(bestCount)

        return k - bestCount[1]


# @leet end
