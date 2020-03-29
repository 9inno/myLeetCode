package leetcode

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct{
		grid [][]int
		result int
	}{
		{[][]int{[]int{1,0,1}, []int{0,0,0}, []int{1,0,1}},2},
	}

	for _ ,test := range tests {
		if actual := MaxDistance(test.grid); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}


