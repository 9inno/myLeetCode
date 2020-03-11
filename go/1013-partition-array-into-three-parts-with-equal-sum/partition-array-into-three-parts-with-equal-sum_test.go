package leetcode

import "testing"


func TestCanThreePartsEqualSum(t *testing.T) {
	tests := [] struct{
		A []int
		result bool
	}{
		{[]int{0,2,1,-6,6,-7,9,1,2,0,1}, true},
		{[]int{0,2,1,-6,6,7,9,-1,2,0,1}, false},
		{[]int{3,3,6,5,-2,2,5,1,-9,4}, true},
	}

	for _ , test := range tests{
		if actual := CanThreePartsEqualSum(test.A); actual != test.result {
			t.Errorf("CanThreePartsEqualSum got %t , expected %t", actual, test.result)
		}

		if actual1 := CanThreePartsEqualSum1(test.A); actual1 != test.result {
			t.Errorf("CanThreePartsEqualSum1 got %t, expected %t", actual1,test.result)
		}
	}

}