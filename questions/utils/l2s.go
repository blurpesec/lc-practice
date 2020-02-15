package utils

type ListNode struct {
  Val int
  Next *ListNode
}

// L2s define
func L2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
