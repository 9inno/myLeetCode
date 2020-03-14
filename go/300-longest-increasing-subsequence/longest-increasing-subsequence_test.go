package leetcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := [] struct {
		nums []int
		result int
	}{
		{[]int {10,9,2,5,3,7,101,18}, 4},
		{[]int {10,9,2,5,3,4}, 3},
	}

	for _ , test := range tests {
		if actual := LengthOfLIS(test.nums); actual != test.result {
			t.Errorf("got %d, expected %d", actual, test.result)
		}
	}


}
