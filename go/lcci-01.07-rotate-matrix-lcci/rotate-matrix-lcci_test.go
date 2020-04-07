package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests :=[] struct{
		matrix [][]int
		result [][]int
	}{
		{[][]int{
			[]int{5, 1, 9,11},
			[]int{2, 4, 8,10},
			[]int{13, 3, 6, 7},
			[]int{15,14,12,16},
		},
		[][]int{
			[]int{15,13, 2, 5},
			[]int{14, 3, 4, 1},
			[]int{12, 6, 8, 9},
			[]int{16, 7,10,11},
		}},
		{
			[][]int{
				[]int{1,2,3},
				[]int{4,5,6},
				[]int{7,8,9},
			},
			[][]int{
				[]int{7,4,1},
				[]int{8,5,2},
				[]int{9,6,3},
			},
		},
	}
	for _ ,test := range tests {
		if Rotate(test.matrix); !reflect.DeepEqual(test.matrix, test.result) {
			t.Errorf("got %v , expected %v", test.matrix, test.result)
		}
	}

}
