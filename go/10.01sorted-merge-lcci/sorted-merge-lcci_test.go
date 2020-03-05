package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T)  {
	tests := []struct{
		A []int
		m int
		B []int
		n int
		result []int
	}{
		{[]int{1,2,3,0,0,0}, 3, []int{2,3,5}, 3, []int{1,2,2,3,3,5} } ,
		{[]int{4,5,7,0,0,0}, 3, []int{7,8,9}, 3, []int{4,5,7,7,8,9} },
		{[]int{4,5,7,10,11,0,0,0,0}, 5, []int{2,3,7,9}, 4, []int{2,3,4,5,7,7,9,10,11} },
	}

	for _ , test := range tests {
		if Merge(test.A, test.m, test.B, test.n); !reflect.DeepEqual(test.A, test.result) {
			t.Errorf(" got %v ; expected %v", test.A, test.result)
		}
	}
}