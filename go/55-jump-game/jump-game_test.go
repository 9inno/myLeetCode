package leetcode

import "testing"

func TestCanJump(t *testing.T) {
	tests:= []struct{
		nums []int
		result bool
	}{
		{[]int{2,3,1,1,4}, true},
		{[]int{3,2,1,0,4}, false},
	}
	for _ , test := range tests {
		if actual := CanJump(test.nums); actual != test.result {
			t.Errorf("got %t, expect %t", actual, test.result)
		}
	}
}