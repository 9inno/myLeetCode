package leetcode

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct{
		head []int
		result []int
	}{
		{[]int{1,2,3,4,5,6}, [] int {2,1,4,3,6,5}},
		{[]int{10,9,8,7,6,5}, [] int {9,10,7,8,5,6}},
	}

	for _, test := range tests {
		if actual := listToSlice(swapPairs(sliceToList(test.head))) ; !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v , expected %v", actual, test.result)
		}

	}
}

func sliceToList(slice []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := head
	for _, value := range slice {
		tmp.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		tmp = tmp.Next
	}
	return head.Next
}

func listToSlice(head *ListNode) []int {
	var slice []int
	for head != nil  {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}