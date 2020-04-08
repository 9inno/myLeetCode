package leetcode

import (
	"testing"
)

var tests []struct{ m, n, k, result int} = []struct{ m, n, k, result int }{{ 2, 3, 1, 3},
	{ 3, 1, 0, 1},
	{ 16, 8, 4, 15},
}

func TestMovingCount(t *testing.T) {

	for _ , test := range tests {
		if actual := MovingCount(test.m, test.n, test.k); actual != test.result {
			t.Errorf("MovingCount got %d ,expected %d", actual, test.result)
		}

	}
}

func TestMovingCount1(t *testing.T) {
	for _ , test := range tests {

		if actual1 := MovingCount1(test.m, test.n, test.k); actual1 != test.result {
			t.Errorf("MovingCount1 got %d ,expected %d", actual1, test.result)
		}
	}
}