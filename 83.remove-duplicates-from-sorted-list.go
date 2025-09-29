// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // Space Complexity = O(1)
 // Time Complexity = O(n)
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		//check if head != next val
		//loop till get to next valid cell
		nextNode := cur.Next

		for nextNode != nil && nextNode.Val == cur.Val {
			nextNode = nextNode.Next
		}

		if nextNode == nil {
			cur.Next = nil
			return head
		} 

		cur.Next = nextNode
		cur = nextNode
	}
	return head
}
// @leet end
