package leetcode

import "testing"

func TestMinIncrementForUnique(t *testing.T) {
	tests := [] struct{
		A []int
		result int
	}{
		{[]int{1,2,2}, 1},
		{[]int{3,2,1,2,1,7}, 6},
	}

	for _ , test := range tests{
		if actual := MinIncrementForUnique(test.A); actual != test.result {
			t.Errorf("got %d , expected %d",actual, test.result)
		}
	}
}