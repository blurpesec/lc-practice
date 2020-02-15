package leetcode

// Thought process:
/*
	Things to consider here:
	1) order
	2) "The new list should be made by splicing together the nodes of the first two lists."

	to implement: 
	1) Select the largest val between the first node from each list. This is root node.
	2) While l1.next != nil check and add l1.next to l2

*/

type ListNode struct {
  Val int
  Next *ListNode
}

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if (l1.Val > l2.Val) {
		
	}

	l2.Next = mergeTwoSortedLists(l1, l2.Next)
	return l2
}