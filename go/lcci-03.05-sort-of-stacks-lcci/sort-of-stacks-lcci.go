package leetcode


type ListNode struct {
	Val int
	Next *ListNode
}

type SortedStack struct {
	stack *ListNode
}


func Constructor() SortedStack {
	return SortedStack{
		stack: nil,
	}
}


func (this *SortedStack) Push(val int)  {
	if this.stack == nil {
		this.stack = &ListNode{
			Val:  val,
			Next: nil,
		}
	} else {
		if this.stack.Val >= val {
			this.stack = &ListNode{
				Val:  val,
				Next: this.stack,
			}
		}else  {
			tmp := this.stack
			for tmp.Next != nil && tmp.Next.Val < val {
				tmp = tmp.Next
			}
			tmp.Next = &ListNode{
				Val:  val,
				Next: tmp.Next,
			}
		}

	}
}


func (this *SortedStack) Pop()  {
	if this.stack != nil {
		this.stack = this.stack.Next
	}
}


func (this *SortedStack) Peek() int {
	if this.stack == nil {
		return -1
	}
	return this.stack.Val
}


func (this *SortedStack) IsEmpty() bool {
	return this.stack == nil
}


/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */