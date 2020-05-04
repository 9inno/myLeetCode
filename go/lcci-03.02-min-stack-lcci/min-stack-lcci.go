package leetcode

import "math"

type MinStack struct {
	stack []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   0,
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.stack) == 0 {
		this.min = x
	} else {
		this.min = Min(this.min, x)
	}
	this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
	index := len(this.stack) - 1
	result := this.stack[index]
	this.stack = this.stack[0 : index]
	if result == this.min {
		tmpMin := math.MaxInt64
		for _ , value := range this.stack {
			tmpMin = Min(tmpMin, value)
		}
		this.min = tmpMin
	}
}


func (this *MinStack) Top() int {
	index := len(this.stack) - 1
	return this.stack[index]
}


func (this *MinStack) GetMin() int {
	return this.min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */