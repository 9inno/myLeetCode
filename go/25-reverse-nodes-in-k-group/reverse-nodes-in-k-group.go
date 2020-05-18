package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	result := &ListNode{Next: head}

	pre := result

	for head != nil {
		tail := pre
		for i := 0; i < k ; i ++ {
			tail = tail.Next
			if tail == nil {
				return result.Next
			}
		}
		next := tail.Next
		head , tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return result.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	tmp := head
	for pre != tail {
		next := tmp.Next
		tmp.Next = pre
		pre = tmp
		tmp = next
	}
	return tail, head
}
