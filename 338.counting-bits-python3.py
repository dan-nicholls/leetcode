# @leet start
class Solution:
    def decimalToBinary(self, n: int) -> str:
        if n == 0:
            return "0"
        binary = ""
        while n > 0:
            binary = str(n % 2) + binary
            n = n // 2
        return binary

    def countBits(self, n: int) -> List[int]:
        result = []
        for i in range(n + 1):
            result.append(self.decimalToBinary(i).count("1"))
        return result


# @leet end
