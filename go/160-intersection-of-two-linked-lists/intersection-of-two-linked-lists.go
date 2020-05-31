package leetcode


type ListNode struct {
 Val int
 Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := map[*ListNode]bool{}
	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _ , ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
