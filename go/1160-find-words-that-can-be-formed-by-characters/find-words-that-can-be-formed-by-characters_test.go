package leetcode

import "testing"

func TestCountCharacters(t *testing.T) {
	tests := []struct{
		words []string
		chars string
		result int
	}{
		{[]string{"cat","bt","hat","tree"}, "atach", 6},
		{[]string{"hello","world","leetcode"}, "welldonehoneyr", 10},

	}

	for _ , test := range tests {
		if actual := CountCharacters(test.words, test.chars); actual != test.result {
			t.Errorf("CountCharacters got %d, expected %d", actual, test.result)
		}

		if actual1 := CountCharacters1(test.words, test.chars); actual1 != test.result {
			t.Errorf("CountCharacters got %d, expected %d", actual1, test.result)
		}
	}
}