package leetcode

import "testing"

func TestLastRemaining(t *testing.T) {
	tests := []struct{
		n,m,result int
	}{
		{5,3, 3},
		{10, 17, 2},
		{5, 1, 4},
		{70866,116922,64165},
	}

	for _ , test := range tests {
		if actual := LastRemaining(test.n, test.m) ; actual != test.result {
			t.Errorf("LastRemaining got %d ,expected %d", actual, test.result)
		}
		if actual1 := LastRemaining1(test.n, test.m) ; actual1 != test.result {
			t.Errorf("LastRemaining1 got %d ,expected %d", actual1, test.result)
		}
	}
}