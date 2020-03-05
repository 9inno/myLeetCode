package leetcode

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {

	tests := [] struct{
		candies int
		numPeople int
		result []int
	}{
		{106, 5, []int {18,21,24,27,16}},
		{10, 3, []int {5,2,3}},
		{7, 4, []int {1,2,3,1}},
	}

	for _, test := range tests{
		if actual := DistributeCandies(test.candies, test.numPeople); !reflect.DeepEqual(actual, test.result) {
			t.Errorf("got %v , expected %v", actual, test.result)
		}
	}
}