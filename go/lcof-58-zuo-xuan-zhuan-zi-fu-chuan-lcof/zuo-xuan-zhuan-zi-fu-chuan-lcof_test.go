package leetcode

import "testing"

func TestReverseLeftWords(t *testing.T) {

	tests := [] struct{
		s string
		n int
		result string
	}{
		{"abcdefg", 2, "cdefgab"},
		{"lrloseumgh", 6, "umghlrlose"},
	}

	for _ , test :=range tests{
		if actual := ReverseLeftWords(test.s,test.n); actual != test.result {
			t.Errorf( " got %s, expected %s", actual, test.result)
		}
	}

}