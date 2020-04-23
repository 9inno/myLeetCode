package leetcode

import "testing"

func TestWaysToChange(t *testing.T) {

	tests := [] struct{
		n int
		result int
	}{
		{5, 2},
		{10, 4},
		{70, 103},
		{0, 1},
		{1000000, 332576607},
	}

	for _, test := range tests {
		if actual := WaysToChange(test.n); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test)
		}
	}
}