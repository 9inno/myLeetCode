package leetcode

import "testing"

func TestFindClosest(t *testing.T) {
	tests := []struct{
		words []string
		word1 string
		word2 string
		result int
	}{
		{[]string{"I","am","a","student","from","a","university","in","a","city"}, "a", "student", 1},
		{[]string{"I","am","a","student","from","a","university","in","a","city"}, "am", "student", 2},
	}

	for _ , test := range tests {
		if actual := FindClosest(test.words, test.word1, test.word2); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}
