package leetcode

import "testing"

func TestOrangesRotting(t *testing.T)  {

	tests := []struct{
		grid [][]int
		result int
	}{
		{[][]int {[] int {2, 1, 1}, [] int {1, 1, 0}, [] int {0, 1, 1} }, 4} ,
		{[][]int {[] int {2, 1, 1}, [] int {0, 1, 1}, []int {1, 0, 1}}, -1},
		{[][] int {[] int {0, 2}}, 0},
	}

	for _ ,test := range tests {
		if actual := OrangesRotting(test.grid); actual != test.result {
			t.Errorf("got %d, expected %d", actual,test.result)
		}
	}
}