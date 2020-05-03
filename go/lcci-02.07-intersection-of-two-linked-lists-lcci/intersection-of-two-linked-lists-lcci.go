package main

func main() {
	a := &ListNode{
		Val:  0,
		Next: nil,
	}
	b := &ListNode{
		Val:  1,
		Next: nil,
	}
	getIntersectionNode(a , b)
}
type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool)
	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
