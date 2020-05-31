package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hashMap := map[*ListNode]bool{}
	for head != nil {
		if _ , ok := hashMap[head]; ok {
			return true
		} else {
			hashMap[head] = true
		}
		head = head.Next
	}
	return false
}
