# @leet start
class Solution:
    def minFallingPathSum(self, grid: List[List[int]]) -> int:
        if grid is None:
            return 0

        n = len(grid)

        for row in range(1, n):
            for col in range(1, n):
                left = grid[row - 1][col - 1] if col > 1 else float("inf")
                up = grid[row - 1][col]
                right = grid[row - 1][col + 1] if col < n - 1 else float("inf")
                grid[row][col] += min(left, up, right)

        return grid[-1]


# @leet end
