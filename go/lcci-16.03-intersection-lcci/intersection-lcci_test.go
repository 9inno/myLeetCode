package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {

	tests := []struct{
		start1 []int
		end1   []int
		start2 []int
		end2   []int
		result []float64
	}{
		{
			[]int{0, 0},
			[]int{1, 0},
			[]int{1, 1},
			[]int{0, -1},
			[]float64{0.5, 0},
		},
		{
			[]int{0, 0},
			[]int{3, 3},
			[]int{1, 1},
			[]int{2, 2},
			[]float64{1, 1},
		},
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{1, 0},
			[]int{2, 1},
			[]float64{},
		},
	}


	for _, test := range tests {
		if actual := Intersection(test.start1, test.end1, test.start2, test.end2); !reflect.DeepEqual(actual,test.result){
			t.Errorf("go %v ,expected %v", actual,test.result)
		}
	}
}