package leetcode

import (
	"reflect"
	"testing"
)

func TestFindContinuousSequence(t *testing.T) {
	tests := []struct{
		target int
		result [][]int
	}{
		{9, [][]int{[]int {2,3,4}, []int {4,5}}},
		{15, [][]int {[]int {1,2,3,4,5}, []int {4,5,6}, []int{7,8}}},
	}

	for _ ,test := range tests{
		if actual := FindContinuousSequence(test.target); !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v , expected %v", actual, test.result)
		}
	}
}