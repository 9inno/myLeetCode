package leetcode

import "testing"

func TestCompressString(t *testing.T) {
	tests := [] struct{
		S string
		result string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abbccd", "abbccd"},
		{"bb", "bb"},
	}

	for _ , test := range tests {
		if actual := CompressString(test.S); actual != test.result {
			t.Errorf("got %s , expected %s", actual, test.result)
		}
	}

}
