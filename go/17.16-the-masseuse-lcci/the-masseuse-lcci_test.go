package leetcode

import "testing"

func TestMassage(t *testing.T) {
	tests := [] struct{
		nums []int
		result int
	}{
		{[]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}, 4173},
		{[]int{1,2,3,1}, 4},
		{[]int{2,7,9,3,1}, 12},
		{[]int{2,1,4,5,3,1,1,3}, 12},
	}

	for _ , test := range tests{
		if actual := Massage(test.nums); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}