package leetcode

import "testing"

func TestWaysToStep(t *testing.T) {
	tests := []struct{
		n int
		result int
	}{
		{100000, 111787461},
		{10000, 126130994},
		{4678, 811388610},
	}


	for _ , test := range tests {
		if actual := WaysToStep(test.n); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}