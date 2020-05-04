package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	hashMap := make(map[*ListNode]bool)
	for head != nil {
		if _ , ok := hashMap[head]; ok {
			return head
		} else {
			hashMap[head] = true
		}
		head = head.Next
	}
	return nil
}
