package leetcode

import (
	"github.com/blurpesec/lc-practice/utils"
)
// Thought process:
/*
	Things to consider here:
	1) carry tens-place && reset when its used

	Steps to complete:
	1) Create an output linked list.
	2) Initialize a carry place to 0
	3) For l1.data and l2.data exist, go through shit
*/


func AddTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	// create output LL
	head := &utils.ListNode{Val: 0, Next: nil}

	// initialize the carry to 0
	carry := 0
	curr := head

	// while l1 or l2 do not equal nil.
	for l1 != nil || l2 != nil {

		// initialize x and y as integers
		var x, y int
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}

		// Set next node
		curr.Next = &utils.ListNode{ Val: (x + y + carry) % 10, Next: nil }
		curr = curr.Next
		carry = (x + y + carry) / 10
		
		// navigate to next steps in input LLs
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// this handles extra carry
	if carry > 0 {
		curr.Next = &utils.ListNode{ Val: carry % 10, Next: nil }
	}

	return head.Next
}