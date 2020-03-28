package leetcode

import "testing"

func TestMinimumLengthEncoding(t *testing.T) {

	tests := []struct{
		words []string
		result int
	}{
		{[]string{"time", "me", "bell"}, 10},
		{[]string{"cat", "dog", "hotdog"},11},

	}

	for _ , test := range tests{
		if actual := MinimumLengthEncoding(test.words); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}