package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	tmp := &ListNode{
		Val:  0,
		Next: head,
	}
	hashMap := map[int]bool{}
	for tmp.Next != nil {
		if _, ok :=hashMap[tmp.Next.Val] ; ok {
			tmp.Next = tmp.Next.Next
		}else {
			hashMap[tmp.Next.Val] = true
			tmp = tmp.Next
		}
	}
	return head
}


func removeDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := head
	hashMap := map[int] bool {head.Val: true}
	for head.Next != nil {
		if _ , ok := hashMap[head.Next.Val]; ok {
			head.Next = head.Next.Next
		} else {
			hashMap[head.Next.Val] = true
			head = head.Next
		}
	}
	return result
}
