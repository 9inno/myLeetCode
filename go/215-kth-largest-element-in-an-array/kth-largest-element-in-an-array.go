package main

import "container/heap"

func main() {
	findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 2)
}
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	for _ ,value := range nums {
		h.Push(value)
	}
	heap.Init(h)
	result := 0

	for i := 0; i < k ;i ++ {
		result = h.Pop().(int)
	}
	return result
}



type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

