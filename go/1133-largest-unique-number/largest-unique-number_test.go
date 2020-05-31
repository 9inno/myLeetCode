package leetcode

import "testing"

func TestLargestUniqueNumber(t *testing.T) {
	tests := []struct{
		nums []int
		result int
	}{
		{[]int{5,7,3,9,4,9,8,3,1}, 8},
		{[]int{9,9,8,8}, -1},
	}

	for _ , test :=range tests {
		if actual := LargestUniqueNumber(test.nums); actual != test.result {
			t.Errorf( "got %d , expected %d", actual, test.result)
		}
	}
}
