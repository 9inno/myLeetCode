package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct{
		height []int
		result int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
	}

	for _ , test := range tests {
		if actual := MaxArea(test.height); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}