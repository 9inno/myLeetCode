package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmpLess := less
	more := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmpMore := more
	for head != nil {
		if head.Val < x {
			tmpLess.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			tmpLess = tmpLess.Next
		} else {
			tmpMore.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			tmpMore = tmpMore.Next
		}
		head = head.Next
	}
	tmpLess.Next = more.Next
	return less.Next
}
