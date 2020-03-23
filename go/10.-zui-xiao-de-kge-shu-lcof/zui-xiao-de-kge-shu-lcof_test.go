package leetcode

import (
	"reflect"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	tests := []struct{
		nums []int
		k int
		result []int
	}{
		{[]int{3,2,1}, 2, []int {1,2}},
		{ [] int {0,1,2,1}, 1 , []int {0}},
	}

	for _ , test := range tests {
		if actual := GetLeastNumbers(test.nums, test.k); !reflect.DeepEqual(actual, test.result){
			t.Errorf(" got %v , expected %v", actual, test.result)
		}
	}
}
