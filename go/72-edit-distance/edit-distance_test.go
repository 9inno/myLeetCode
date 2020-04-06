package leetcode

import "testing"

func TestMinDistance(t *testing.T) {

	tests := []struct{
		word1 string
		word2 string
		result int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},

	}

	for _ , test := range tests {
		if actual := MinDistance(test.word1, test.word2); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}