# @leet start
class Solution:
    # Attempted 2 - Stack & Reverse Iteration
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        if len(temperatures) == 1:
            return [0]

        result = [0] * len(temperatures)

        stack = []

        for currentDayIndex in range(len(temperatures) - 1, -1, -1):
            print(currentDayIndex)
            while stack and temperatures[currentDayIndex] >= temperatures[stack[-1]]:
                stack.pop()
            if stack:
                result[currentDayIndex] = stack[-1] - currentDayIndex
            stack.append(currentDayIndex)
        return result

    # Attempted 1 - Loops
    # Doesn't handle large input lengths well
    def dailyTemperatures1(self, temperatures: List[int]) -> List[int]:
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
