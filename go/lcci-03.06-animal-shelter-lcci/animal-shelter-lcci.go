package leetcode

import "container/list"

type AnimalShelf struct {
	dogList *list.List
	catList *list.List
	seq *list.List
}


func Constructor() AnimalShelf {
	return AnimalShelf{
		dogList: list.New(),
		catList: list.New(),
		seq:     list.New(),
	}
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	var tmp *list.List
	if animal[1] == 0 {
		tmp = this.catList
	} else {
		tmp = this.dogList
	}
	element := this.seq.PushBack(animal)
	tmp.PushBack(element)
}


func (this *AnimalShelf) DequeueAny() []int {
	if this.seq.Len() == 0 {
		return []int{-1, -1}
	}
	animal := this.seq.Remove(this.seq.Front()).([]int)
	var tmp *list.List
	if animal[1] == 0 {
		tmp = this.catList
	} else {
		tmp = this.dogList
	}
	tmp.Remove(tmp.Front())
	return animal
}


func (this *AnimalShelf) DequeueDog() []int {
	if this.dogList.Len() == 0 {
		return []int{-1, -1}
	}
	element := this.dogList.Remove(this.dogList.Front()).(*list.Element)
	return this.seq.Remove(element).([]int)
}


func (this *AnimalShelf) DequeueCat() []int {
	if this.catList.Len() == 0 {
		return []int{-1, -1}
	}
	element := this.catList.Remove(this.catList.Front()).(*list.Element)
	return this.seq.Remove(element).([]int)
}


/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */