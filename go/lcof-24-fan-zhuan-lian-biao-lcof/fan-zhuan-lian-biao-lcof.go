package main

func main() {
	a := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},

				},
			},
		},
	}
	reverseList(a)

}
type ListNode struct {
 Val int
 Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre

}
