package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	tests := [] struct{
		coin []int
		amount int
		result int
	}{
		{[]int {1, 2, 5}, 11 , 3},
		{[]int{2}, 3, -1},
	}

	for _, test := range tests{
		if actual := CoinChange(test.coin, test.amount); actual != test.result {
			t.Errorf("got %d, expected %d", actual, test.result)
		}
	}
}