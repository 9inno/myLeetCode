package leetcode

type TripleInOne struct {
	stack [][]int
	max int
}


func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack: [][]int{[]int{}, []int{}, []int{}},
		max: stackSize,
	}
}


func (this *TripleInOne) Push(stackNum int, value int)  {
	if len(this.stack[stackNum]) < this.max {
		this.stack[stackNum] = append(this.stack[stackNum], value)
	}
}


func (this *TripleInOne) Pop(stackNum int) int {
	if len(this.stack[stackNum]) == 0 {
		return -1
	}
	index := len(this.stack[stackNum]) - 1
	result := this.stack[stackNum][index]
	this.stack[stackNum] = this.stack[stackNum][0: index]
	return result

}


func (this *TripleInOne) Peek(stackNum int) int {
	if len(this.stack[stackNum]) == 0 {
		return -1
	}
	index := len(this.stack[stackNum]) - 1
	return this.stack[stackNum][index]
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.stack[stackNum]) == 0
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
