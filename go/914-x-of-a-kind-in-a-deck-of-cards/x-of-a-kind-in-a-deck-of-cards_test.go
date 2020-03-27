package leetcode

import "testing"

func TestHasGroupsSizeX(t *testing.T) {

	tests := [] struct{
		deck []int
		result bool
	}{
		{[]int{1,2,3,4,4,3,2,1}, true},
		{[]int{1,1,1,2,2,2,3,3}, false},
		{[]int{1}, false},
		{[]int{1,1},true},
		{[]int{1,1,2,2,2,2},true},
	}

	for _ , test := range tests{
		if actual := HasGroupsSizeX(test.deck); actual != test.result {
			t.Errorf( "got %t , expected %t", actual, test.result)
		}
	}
}
