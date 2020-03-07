package leetcode

import (
	"testing"
)

func TestQueue(t *testing.T) {

	queue := Constructor()
	queue.Push_back(1)
	queue.Push_back(2)
	if result:= queue.maxValue; result != 2 {
		t.Errorf("got %d, expected %d", result, 2)
	}
	if result := queue.Pop_front() ;  result!= 1 {
		t.Errorf("got %d, expected %d", result, 1)
	}
	if result:= queue.maxValue; result != 2 {
		t.Errorf("got %d, expected %d", result, 2)
	}

}


