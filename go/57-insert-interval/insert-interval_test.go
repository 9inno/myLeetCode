package leetcode

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct{
		intervals [][]int
		newIntervals []int
		result [][]int
	}{
		{[][]int{[]int{1,3}, []int{6,9}}, []int{2,5}, [][]int{[]int{1,5},[]int{6,9}}},
		{[][]int{[]int{1,2}, []int{3,5}, []int{6,7}, []int{8,10}, []int{12,16}}, []int{4,8}, [][]int{[]int{1,2},[]int{3,10}, []int{12,16}}},
	}

	for _ , test :=range tests {
		if actual := Insert(test.intervals, test.newIntervals) ; !reflect.DeepEqual(actual, test.result) {
			t.Errorf( "got %v , expected %v ", actual, test.result)
		}
	}
}