package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests:= [] struct{
		prices []int
		result int
	}{
		{[]int {7,1,5,3,6,4}, 5},
		{[]int {7,6,4,3,1}, 0},
	}

	for _ ,test := range tests{
		if actual := MaxProfit(test.prices); actual != test.result {
			t.Errorf("got %d, expected %d", actual, test.result)
		}
	}
}