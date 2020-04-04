package leetcode

import "testing"

func TestTrap(t *testing.T) {

	tests := []struct{
		height []int
		result int
	}{
		{[]int{0,1,0,2,1,0,1,3,2,1,2,1}, 6},
	}

	for _ , test := range tests {
		if actual := Trap(test.height); actual != test.result {
			t.Errorf("got %d ,expected %d", actual, test.result)
		}
	}
}
