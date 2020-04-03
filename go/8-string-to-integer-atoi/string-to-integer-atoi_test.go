package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct{
		str string
		result int
	}{
		{"   -42", -42},
		{"words and 987", 0},
		{"-91283472332",  -2147483648},
		{"  0000000000012345678", 12345678},
	}

	for _ , test := range tests {
		if actual := MyAtoi(test.str); actual  != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}