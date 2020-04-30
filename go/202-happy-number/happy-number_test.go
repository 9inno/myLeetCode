package leetcode

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct{
		n int
		result bool
	}{
		{19, true},
		{20, false},
		{0, false},
		{12345,false},
	}

	for _ , test :=range tests {
		if actual := IsHappy(test.n); actual != test.result {
			t.Errorf("got %t , expected %t", actual, test.result)
		}
	}
}