# @leet start
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        total = 0

        while mainTank >= 5:
            mainTank -= 5
            total += 50

            if additionalTank > 0:
                additionalTank -= 1
                mainTank += 1
        total += mainTank * 10
        return total


# @leet end
