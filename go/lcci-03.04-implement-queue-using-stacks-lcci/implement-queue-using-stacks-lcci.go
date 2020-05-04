package leetcode

type MyQueue struct {
	queue []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		queue: []int{},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.queue) > 0 {
		result := this.queue[0]
		this.queue = this.queue[1:]
		return result
	}
	return -1
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}
	return -1
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
