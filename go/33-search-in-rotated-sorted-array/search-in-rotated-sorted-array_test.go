package leetcode

import "testing"

func TestSearch(t *testing.T) {
	tests := [] struct{
		nums []int
		target int
		result int
	}{
		{[]int{4,5,6,7,9,11,32,54,66,77,0,1,2}, 32, 6},
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{4,5,6,7,0,1,2}, 3, -1},
	}

	for _ , test := range tests {
		if actual := Search(test.nums, test.target); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}

}