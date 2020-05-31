package leetcode

import "testing"

func TestGuessNumber(t *testing.T) {

	tests := []struct{
		n int
		target int
		result int
	}{
		{100, 67, 67},
		{10, 6, 6},
	}

	for _ ,test :=range tests {
		Target = test.target
		if actual := GuessNumber(test.n); actual !=test.result {
			t.Errorf("got %d ,expected %d", actual, test.result)
		}
	}
}
