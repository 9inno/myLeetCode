package leetcode

import "testing"

func TestReverseWords(t *testing.T) {

	tests := []struct{
		s string
		result string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for _ , test := range tests {
		if actual := ReverseWords(test.s); actual != test.result {
			t.Errorf("got %s , expected %s", actual, test.result)
		}
	}
}