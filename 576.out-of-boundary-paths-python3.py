# @leet start
class Solution:
    def findPaths(
        self, m: int, n: int, maxMove: int, startRow: int, startColumn: int
    ) -> int:
        if maxMove == 0:
            return 0

        outOfBoundsPaths = 0

        availableMoves = [
            (startRow - 1, startColumn),
            (startRow + 1, startColumn),
            (startRow, startColumn - 1),
            (startRow, startColumn + 1),
        ]

        print(f"availableMoves: {availableMoves}")

        for moveRow, moveCol in availableMoves:
            print(f"({startRow}, {startColumn}) -> ({moveRow}, {moveCol})")
            if moveRow < 0 or moveRow > m or moveCol < 0 or moveCol > n:
                outOfBoundsPaths += 1
            else:
                outOfBoundsPaths += self.findPaths(m, n, maxMove - 1, moveRow, moveCol)

        return outOfBoundsPaths


# @leet end
