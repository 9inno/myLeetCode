package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {

	tests := []struct{
		list []int
		result []int
	}{
		{[]int{1,2,3,4,5,6}, [] int {6,5,4,3,2,1}},
		{[]int{10,9,8,7,6,5}, [] int {5,6,7,8,9,10}},
	}
	for _ , test := range tests{
		if actual := listToSlice(ReverseList(sliceToList(test.list))); !reflect.DeepEqual(actual, test.result) {
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