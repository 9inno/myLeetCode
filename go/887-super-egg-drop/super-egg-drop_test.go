package leetcode

import "testing"

var tests = [] struct{
	K int
	N int
	result int
}{
	{1, 2, 2},
	{2, 6, 3},
	{3, 14, 4},
}

func TestSuperEggDrop(t *testing.T) {

	for _ , test := range tests {
		if actual := SuperEggDrop(test.K, test.N); actual != test.result {
			t.Errorf( "got %d , expected %d", actual, test.result)
		}
	}
}

func TestSuperEggDrop1(t *testing.T) {
	for _ , test := range tests {
		if actual := SuperEggDrop1(test.K, test.N); actual != test.result {
			t.Errorf( "got %d , expected %d", actual, test.result)
		}
	}
}