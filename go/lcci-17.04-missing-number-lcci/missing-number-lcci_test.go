package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {


	tests := []struct {
		nums []int
		result int
	}{
		{[]int{3,0,1}, 2},
		{[]int{0,1,2,3}, 4},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},
	}

	for _ , test := range tests {
		if actual := MissingNumber(test.nums); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}
