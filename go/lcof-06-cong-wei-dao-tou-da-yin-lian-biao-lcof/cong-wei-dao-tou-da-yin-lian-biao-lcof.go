package leetcode

type ListNode struct {
 Val int
 Next *ListNode
}

func reversePrint(head *ListNode) []int {

	result := []int{}

	for head != nil {
		result = append([]int{head.Val}, result...)
		head = head.Next
	}
	return result
}
