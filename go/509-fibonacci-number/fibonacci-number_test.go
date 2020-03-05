package leetcode

import "testing"

func TestFib(t *testing.T) {
	tests := []struct{N , Result int} {
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _ , test := range tests {
		if actual := Fib(test.N); actual != test.Result {
			t.Errorf("Fib(%d) got %d expected %d", test.N, actual, test.Result)
		}
	}
}