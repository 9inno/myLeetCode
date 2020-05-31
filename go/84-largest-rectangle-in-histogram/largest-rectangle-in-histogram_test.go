package leetcode

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct{
		nums []int
		result int
	}{
		{[]int{2,1,5,6,2,3}, 10},
	}

	for _ , test := range tests {
		if actual := LargestRectangleArea(test.nums); actual != test.result {
			t.Errorf("got %d , expected %d", actual ,test.result)
		}
	}
}
