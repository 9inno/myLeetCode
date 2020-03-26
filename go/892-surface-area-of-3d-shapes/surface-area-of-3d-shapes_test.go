package leetcode

import "testing"

func TestSurfaceArea(t *testing.T) {
	tests := []struct{
		grid [][]int
		result int
	}{
		{[][]int{[]int{2}}, 10},
		{[][]int{[]int{1,2}, []int{3,4}}, 34},
		{[][]int{[]int{1,0}, []int{0,2}},16},
		{[][]int{[]int{1,1,1}, []int{1,0,1}, []int{1,1,1}}, 32},
		{[][]int{[]int{2,2,2}, []int{2,1,2}, []int{2,1,2}},46},

	}

	for _ , test := range tests{
		if actual := SurfaceArea(test.grid); actual != test.result {
			t.Errorf("got %d, expected %d", actual, test.result)
		}
	}

}