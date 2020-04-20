package leetcode

import "testing"


var tests = []struct{
	grid [][]byte
	result int
}{
	{[][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	},1},
	{[][]byte{
		[]byte("11000"),
		[]byte("11010"),
		[]byte("00100"),
		[]byte("00011"),
	},4},
	{[][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	},3},
}

func TestNumIslands(t *testing.T) {
	for _ , test := range tests {
		if actual := NumIslands(test.grid) ; actual != test.result {
			t.Errorf("got %d ,expected %d", actual, test.result)
		}
	}
}

func TestNumIslands1(t *testing.T) {
	for _ , test := range tests {
		if actual := NumIslands1(test.grid) ; actual != test.result {
			t.Errorf("got %d ,expected %d", actual, test.result)
		}
	}
}
