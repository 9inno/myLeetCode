package leetcode

import "testing"

func TestDecodeString(t *testing.T) {

	tests := []struct{
		s string
		result string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, test := range tests {
		if actual := DecodeString(test.s); actual !=test.result {
			t.Errorf("got %s , expected %s", actual, test.result)
		}
	}
}