package leetcode

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct{
		s string
		result bool
	}{
		{"abc", true},
		{"aab", false},
		{"leetcode", false},
	}

	for _, test := range tests {
		if actual := IsUnique(test.s); actual != test.result {
			t.Errorf("got %t, expected %t", actual, test.result)
		}
	}
}
