package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	tests := [] struct{
		s string
		result int
	}{
		{"asdqweadarqwezxfcasdqezxczgveasfasqweaxczxxcqaseqwascshdszhvasfcfarvvgabfasvcacfxgacvghhvczsfqahavgvaesrcfaxewrx", 11},
		{"pwwkew", 3},
		{"bbbb", 1},
		{"abcabcbb", 3},
	}

	for _ ,test := range tests {
		if actual := LengthOfLongestSubstring(test.s); actual != test.result {
			t.Errorf("got %d, expected %d", actual, test.result)
		}
	}
}