package leetcode
type ListNode struct {
     Val int
     Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = head.Next.Next
	tmp.Next = head
	tmp.Next.Next = swapPairs(tmp.Next.Next)
	return tmp
}