package leetcode


type CustomStack struct {
	maxSize int
	stack []int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		stack:   []int{},
	}
}


func (this *CustomStack) Push(x int)  {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}


func (this *CustomStack) Pop() int {
	if len(this.stack) > 0 {
		result := this.stack[len(this.stack) - 1];
		this.stack = this.stack[0 : len(this.stack) - 1]
		return result
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {

	var length int
	if k > len(this.stack) {
		length = len(this.stack)
	} else {
		length = k
	}
	for i := 0; i < length; i++ {
		this.stack[i] = this.stack[i] + val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */