package leetcode

import "testing"

func TestSubarraysDivByK(t *testing.T) {

	tests := []struct{
		nums []int
		k int
 		result int
	}{
		{[]int{4,5,0,-2,-3,1}, 5, 7},
		{[]int{}, 10 , 0},
	}

	for _ , test := range tests {
		if actual := SubarraysDivByK(test.nums, test.k); actual != test.result {
			t.Errorf("got %d , expected %d" , actual , test.result)
		}
	}
}