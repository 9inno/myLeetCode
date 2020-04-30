package leetcode

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := []struct{
		s1 string
		s2 string
		result bool
	}{
		{"asd", "dsa", true},
		{"asd", "zxc", false},
		{"dd", "bb", false},
	}

	for _ , test :=range tests {
		if actual := CheckPermutation(test.s1, test.s2); actual != test.result {
			t.Errorf("got %t, expected %t", actual ,test.result)
		}
	}
}