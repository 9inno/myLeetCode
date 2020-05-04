package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast := head
	slow := head
	for fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	reverseListNode :=reverse(slow.Next)
	for reverseListNode != nil {
		if reverseListNode.Val != head.Val {
			return false
		}
		reverseListNode = reverseListNode.Next
		head = head.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	current := head
	var pre *ListNode
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}