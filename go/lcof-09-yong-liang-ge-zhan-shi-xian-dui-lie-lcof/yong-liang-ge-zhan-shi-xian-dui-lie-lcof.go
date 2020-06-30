package leetcode

import "container/list"

type CQueue struct {
	l1, l2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		l1: list.New(),
		l2: list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.l1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	if this.l2.Len() == 0 {
		for this.l1.Len() > 0 {
			this.l2.PushBack(this.l1.Remove(this.l1.Back()))
		}
	}
	if this.l2.Len() != 0 {
		e := this.l2.Back()
		this.l2.Remove(e)
		return e.Value.(int)
	}
	return -1

}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
