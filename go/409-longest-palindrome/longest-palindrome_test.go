package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {

	tests := []struct{
		s string
		result int
	}{
		{"abccccdd", 7},
		{"asdasdasdqweqweqweASDZXCKLQWECNMAQWEqqqewwwddddsssszxxxddaa", 47},
	}

	for _ , test := range tests{
		if actual := LongestPalindrome(test.s); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}