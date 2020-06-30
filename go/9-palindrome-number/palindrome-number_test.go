package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		x int
		result bool
	}{
		{121, true},
		{0, true},
		{-121, false},
		{10, false},
	}


	for _ , test := range tests {
		if actual := IsPalindrome(test.x); actual != test.result {
			t.Errorf(" IsPalindrome got %t, expected %t", actual, test.result)
		}
		if actual := IsPalindrome1(test.x); actual != test.result {
			t.Errorf(" IsPalindrome1 got %t, expected %t", actual, test.result)
		}
	}
}
