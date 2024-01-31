# @leet start
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        if len(temperatures) == 1:
            return [0]

        result = [0] * len(temperatures)

        for currentDayIndex in range(0, len(temperatures) - 1):
            for compareIndex in range(currentDayIndex + 1, len(temperatures)):
                if temperatures[currentDayIndex] < temperatures[compareIndex]:
                    result[currentDayIndex] = compareIndex - currentDayIndex
                    break

        return result


# @leet end
