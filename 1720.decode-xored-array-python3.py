# @leet start
class Solution:
    def decode(self, encoded: List[int], first: int) -> List[int]:
        # 6 = 110, 2 = 010  XOR = 100 = 4
        result = [first]
        for val in encoded:
            result.append(val ^ result[-1])
        print(result)
        return result

    #


# @leet end
