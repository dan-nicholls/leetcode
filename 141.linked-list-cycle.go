// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]struct{})
	
	cur := head
	
	for {
		if cur == nil {
			return false
		}
		if _, ok := visited[cur]; ok {
			return true
		}

		visited[cur] = struct{}{}

		cur = cur.Next
	}
}
// @leet end
