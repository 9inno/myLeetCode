package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {

	tests := [] struct{
		numberRow int
		result [][]int
	}{
		{5 , [][]int{[]int {1}, []int{1,1}, []int{1,2,1} , []int {1,3,3,1}, []int {1,4,6,4,1}}},
		{2 , [][]int{[]int {1}, []int{1,1}}},
		{0 , [][]int{}},
	}

	for _ ,test := range tests{
		if actual := Generate(test.numberRow); !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v, expected %v", actual, test.result)
		}
	}
}