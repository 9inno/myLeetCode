package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmpListNode := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmpListNode
}