package leetcode

import (
	"github.com/blurpesec/lc-practice/questions/utils"
)
// Thought process:
/*
	Things to consider here:
	1) order
	2) "The new list should be made by splicing together the nodes of the first two lists."

	to implement: 
	1) Select the largest val between the first node from each list. This is root node.
	2) While l1.next != nil check and add l1.next to l2

*/


func mergeTwoSortedLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	// If we've reached the end of l1, return l2
	if l1 == nil {
		return l2
	}

	// If we've reached the end of l2, return l1
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		// recursively merge the two lists, starting with l1
		l1.Next = mergeTwoSortedLists(l1.Next, l2)
		return l1
	}
	// recursively merge the two lists, starting with l2
	l2.Next = mergeTwoSortedLists(l1, l2.Next)
	return l2
}