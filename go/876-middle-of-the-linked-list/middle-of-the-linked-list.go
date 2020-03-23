package leetcode



type ListNode struct {
    Val int
    Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	result := []*ListNode{}

	for head != nil {
		result = append(result, head)
		head = head.Next
	}
	p := len(result) / 2
	return result[p]

}


func MiddleNode1(head *ListNode) *ListNode {

	fast := head
	for fast != nil && fast.Next != nil  {
		head = head.Next
		fast = fast.Next.Next
	}
	return head
}