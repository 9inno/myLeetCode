package leetcode

type MaxQueue struct {
	maxValue int
	queue []int
}


func Constructor() MaxQueue {

	return MaxQueue{
		maxValue: -1,
		queue:    []int{},
	}
	
}


func (this *MaxQueue) Max_value() int {
	return this.maxValue
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	if value > this.maxValue {
		this.maxValue = value
	}

}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	result := this.queue[0]
	this.queue = this.queue[1:]

	if result == this.maxValue {
		tmpMax := -1
		for _ , value := range this.queue {
			if tmpMax < value {
				tmpMax = value
			}
		}
		this.maxValue = tmpMax
	}

	return result

}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
