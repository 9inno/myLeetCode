package leetcode

import "testing"

func TestRob(t *testing.T) {
	tests := []struct{
		nums []int
		result int
	}{
		{[]int{1,2,3,1}, 4},
		{[]int{2,7,9,3,1}, 12},
	}

	for _ , test := range tests {
		if actual := Rob(test.nums); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}
