package leetcode

import "testing"

func TestCanMeasureWater(t *testing.T) {
	tests := []struct{
		x int
		y int
		z int
		result bool
	} {
		{3, 5, 4, true},
		{2, 6, 5, false},
	}

	for _ , test := range tests {
		if actual := CanMeasureWater(test.x, test.y, test.z); actual != test.result {
			t.Errorf("got %t ,expected %t", actual, test.result)
		}
	}
}