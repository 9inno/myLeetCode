package leetcode

import "testing"

func TestIsRectangleOverlap(t *testing.T) {
	tests := [] struct{
		rec1 []int
		rec2 []int
		result bool
	}{
		{[]int{0,0,2,2}, []int{1,1,3,3}, true},
		{[]int{0,0,1,1}, []int{1,0,2,1}, false},
	}

	for _ , test := range tests {
		if actual := IsRectangleOverlap(test.rec1, test.rec2); actual != test.result {
			t.Errorf("IsRectangleOverlap got %t, expected %t", actual,test.result)
		}
		if actual1 := IsRectangleOverlap1(test.rec1, test.rec2); actual1 != test.result {
			t.Errorf("IsRectangleOverlap1 got %t, expected %t", actual1,test.result)
		}
	}
}