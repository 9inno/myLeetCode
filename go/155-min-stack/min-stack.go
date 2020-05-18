package leetcode

import "math"

type MinStack struct {
	stack []int
	minValue int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minValue: math.MaxInt64,
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	this.minValue = min(this.minValue, x)
}


func (this *MinStack) Pop()  {
	index := len(this.stack) - 1
	tmpValue := this.stack[index]
	this.stack = this.stack[:index]
	if tmpValue == this.minValue {
		tmpMin := math.MaxInt64
		for _ , value := range this.stack {
			tmpMin = min(tmpMin, value)
		}
		this.minValue = tmpMin
	}
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.minValue
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
