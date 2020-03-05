package leetcode

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := [] struct{
		index int
		result []int
	}{
		{3, []int {1,3,3,1}},
		{0, []int {1}},
		{6, []int {1,6,15,20,15,6,1}},
		{9, []int {1,9,36,84,126,126,84,36,9,1}},
	}

	for _ , test := range tests{
		if actual := GetRow(test.index); !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v , expected %v", actual,test.result)
		}
	}
}