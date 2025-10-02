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
func removeNthFromEndSinglePass(head *ListNode, n int) *ListNode {
	pointA := head
	var pointB *ListNode 
	var nodeBefore *ListNode
	var nodeAfter *ListNode
	count := 0

	for pointA != nil {
		//fmt.Printf("A: %+v\n", pointA) 
		//fmt.Printf("B: %+v\n", pointB) 
		count++
		
		pointA = pointA.Next
		if count == n {
			pointB = head
			nodeAfter = head.Next
		}
		if count > n {
			nodeBefore = pointB
			pointB = pointB.Next
			nodeAfter = pointB.Next
		}
	}

	//fmt.Printf("A: %+v\n", pointA) 
	//fmt.Printf("B: %+v\n", pointB) 
	//fmt.Printf("Before: %+v\n", nodeBefore) 
	//fmt.Printf("After: %+v\n", nodeAfter) 

	if nodeBefore != nil {
		nodeBefore.Next = nodeAfter
	}
	if pointB == head {
		return nodeAfter
	}
	return head
}

// Space Complexity = O(n)
// Time Complexity = O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	cur := head
	seen := make([]*ListNode, 0)

	for cur != nil {
		seen = append(seen, cur)
		cur = cur.Next
	}

	if (len(seen) == 1 && n == 1) {
		return nil
	}

	var nodeBefore *ListNode
	var nodeAfter *ListNode

	if n != 1 {
		nodeAfter = seen[len(seen)-n+1]
	}
	if n != len(seen) {
		nodeBefore = seen[len(seen)-n-1]
		nodeBefore.Next = nodeAfter
	}

	//fmt.Printf("Before: %+v\n", nodeBefore) 
	//fmt.Printf("After: %+v\n", nodeAfter) 

	// Check if current head is being removed
	if n == len(seen) && nodeAfter != nil {
		head = nodeAfter
	}

	return head
}
// @leet end
