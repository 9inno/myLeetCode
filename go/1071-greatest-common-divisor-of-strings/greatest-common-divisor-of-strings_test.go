package leetcode

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct{
		str1 string
		str2 string
		result string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "AB", "AB"},
		{"ABABABAB", "ABAB", "ABAB"},
		{"LEET", "CODE", ""},
	}

	for _ ,test := range tests {
		if actual := GcdOfStrings(test.str1, test.str2); actual != test.result {
			t.Errorf("got %s, expected %s", actual, test.result)
		}
	}
}
