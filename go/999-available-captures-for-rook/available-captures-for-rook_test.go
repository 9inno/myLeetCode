package leetcode

import "testing"

func TestNumRookCaptures(t *testing.T) {
	 tests := [] struct{
	 	board [][]byte
	 	result int
	 }{
	 	{[][]byte{[]byte("........"),
	 		[]byte("...p...."),
			[]byte("..pp...."),
			[]byte(".ppR.p.p"),
			[]byte("...p...."),
			[]byte(".....p.."),
			[]byte("...p...."),
			[]byte("........")}, 4},
	 }

	 for _ ,test :=range tests{
	 	if actual := NumRookCaptures(test.board); actual != test.result {
	 		t.Errorf("got %d , expected %d", actual, test.result)
		}
	 }
}
