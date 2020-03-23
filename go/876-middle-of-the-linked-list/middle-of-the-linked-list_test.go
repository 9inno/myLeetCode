package leetcode

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := [] struct{
		nodeList []int
		result []int
	}{
		{[]int{1,2,3,4,5}, []int{3,4,5}},
		{[]int{1,2,3,4,5,6}, []int{4,5,6}},
	}

	for _, test := range tests{
		if actual := listToNode(MiddleNode(sliceToList(test.nodeList)));  !reflect.DeepEqual(actual, test.result){
			t.Errorf("MiddleNode got %v, expected %v", actual, test.result)
		}
		if actual := listToNode(MiddleNode1(sliceToList(test.nodeList)));  !reflect.DeepEqual(actual, test.result){
			t.Errorf("MiddleNode1 got %v, expected %v", actual, test.result)
		}
	}
}

func sliceToList(slice []int) *ListNode {
	head := ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := &head
	for _ , value := range slice{
		tmp.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		tmp = tmp.Next
	}
	return head.Next
}

func listToNode(node *ListNode) []int {
	var slice []int
	for node !=nil  {
		slice = append(slice, node.Val)
		node = node.Next
	}
	return slice
}