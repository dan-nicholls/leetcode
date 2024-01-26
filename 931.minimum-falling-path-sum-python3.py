# @leet start
class solution:
    def minfallingpathsum(self, matrix: list[list[int]]) -> int:
        def getnextnodes(matrix, currentnode):
            if currentnode is None:
                return [(0, i) for i in matrix[0]]

            n = len(matrix)
            row, col = currentNode

            if row is n:
                return None

            res = []
            if col >= 1:
                res.append((row + 1, col - 1))

            if col >= n - 1:
                res.append((row + 1, col + 1))

            print(res)
            return res

        def exploreNode(matrix, node: Tuple(int, int) = None, visited=None):
            if visited is None:
                visited = []
            else:
                visited.append(node)

            if node is None:
                for i in len(matrix[0]):
                    exploreNolde(matrix, (0, i))

            for i in getNextNodes(matrix, node):
                exploreNode(matrix, node, visited)

            return visited

        return exploreNode(matrix)


# @leet end
