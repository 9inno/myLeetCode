package leetcode

type StackOfPlates struct {
	stacks  [][]int
	cap     int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stacks: make([][]int, 0),
		cap: cap,
	}
}

func (this *StackOfPlates) Push(val int)  {
	if this.cap == 0 {
		return
	}
	i := len(this.stacks) - 1
	if i == -1 {
		this.stacks = append(this.stacks, []int{val})
		return
	}
	if len(this.stacks[i]) < this.cap {
		this.stacks[i] = append(this.stacks[i], val)
	} else {
		this.stacks = append(this.stacks, []int{val})
	}
}


func (this *StackOfPlates) Pop() int {
	return this.PopAt(len(this.stacks) - 1)
}


func (this *StackOfPlates) PopAt(index int) int {
	if index == -1 || len(this.stacks) <= index {
		return -1
	}
	len_ := len(this.stacks[index])
	ret := this.stacks[index][len_-1]
	if len_ == 1 {
		this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
	} else {
		this.stacks[index] = this.stacks[index][:len_-1]
	}
	return ret
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
