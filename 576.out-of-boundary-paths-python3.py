# @leet start
class Solution:
    def findPaths(
        self, m: int, n: int, maxMove: int, startRow: int, startColumn: int
    ) -> int:
        MODULO = 10**9 + 7
        moveCache = {}

        def calcMoves(currentRow: int, currentCol: int, remainingMoves: int) -> int:
            if (currentRow, currentCol, remainingMoves) in moveCache:
                return moveCache[(currentRow, currentCol, remainingMoves)]

            if currentRow < 0 or currentRow >= m or currentCol < 0 or currentCol >= n:
                return 1

            if remainingMoves == 0:
                return 0

            outOfBoundsPaths = 0

            availableMoves = [
                (currentRow - 1, currentCol),
                (currentRow + 1, currentCol),
                (currentRow, currentCol - 1),
                (currentRow, currentCol + 1),
            ]

            print(f"availableMoves: {availableMoves}, remainingMoves: {remainingMoves}")

            for moveRow, moveCol in availableMoves:
                outOfBoundsPaths += calcMoves(moveRow, moveCol, remainingMoves - 1)
                outOfBoundsPaths %= MODULO

            moveCache[(currentRow, currentCol, remainingMoves)] = outOfBoundsPaths
            return outOfBoundsPaths

        return calcMoves(startRow, startColumn, maxMove)


# @leet end
