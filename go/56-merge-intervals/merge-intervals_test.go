package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := [] struct{
		intervals [][]int
		result [][]int
	} {
		{
			[][]int{
				[]int{1,3},
				[]int{2,6},
				[]int{8,10},
				[]int{15,18},
			},
			[][]int{
				[]int{1,6},
				[]int{8,10},
				[]int{15,18},
			},

		},
		{
			[][]int{
				[]int{1,4},
				[]int{4,5},
			},
			[][]int{
				[]int{1,5},
			},
		},
	}

	for _ ,test := range tests{
		if actual := Merge(test.intervals); !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v , expected %v", actual, test.result)
		}
	}
}