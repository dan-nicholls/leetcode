# @leet start
class Solution:
    def dayOfYear(self, date: str) -> int:
        year, month, day = map(int, date.split("-"))
        daysInMonths = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

        # Check Leap Year
        if (year % 4 == 0 and year % 100 != 0) or (year % 400 == 0):
            daysInMonths[1] = 29

        totalDays = day

        # Completed Months
        completedMonths = max(0, month - 1)
        print(f"Completed Months: {completedMonths}")

        for i in range(completedMonths):
            print(i)
            totalDays += daysInMonths[i]

        return totalDays


# @leet end
