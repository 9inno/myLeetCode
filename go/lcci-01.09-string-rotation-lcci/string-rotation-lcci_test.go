package leetcode

import "testing"

func TestIsFlipedString(t *testing.T) {

	tests := []struct{
		s1 string
		s2 string
		result bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"aaaa", "aaaa", true},
		{"aa", "aba" , false},
		{"", "" , true},
	}

	for _ ,test := range tests {
		if actual := IsFlipedString(test.s1, test.s2); actual != test.result {
			t.Errorf("got %t, expected %t", actual, test.result)
		}
	}
}