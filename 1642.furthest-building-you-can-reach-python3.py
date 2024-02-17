# @leet start
class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        heightDifHeap = []

        for i in range(len(heights) - 1):
            heightDif = heights[i + 1] - heights[i]

            if heightDif <= 0:
                continue

            heapq.heappush(heightDifHeap, heightDif)

            if len(heightDifHeap) > ladders:
                brick_need = heapq.heappop(heightDifHeap)
                bricks -= brick_need

                if bricks < 0:
                    return i

        return len(heights) - 1


# @leet end
