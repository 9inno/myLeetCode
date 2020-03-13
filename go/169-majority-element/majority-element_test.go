package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := [] struct{
		nums []int
		result int
	}{
		{[]int{3,2,3} , 3 },
		{[]int{1}, 1},
		{[]int{2,2,1,1,1,2,2}, 2},
	}

	for _ , test := range tests {
		if actual := MajorityElement(test.nums) ; actual != test.result {
			t.Errorf("MajorityElement got %d , expected %d", actual, test.result)
		}
		if actual := MajorityElement1(test.nums) ; actual != test.result {
			t.Errorf("MajorityElement1 got %d , expected %d", actual, test.result)
		}
	}
}
