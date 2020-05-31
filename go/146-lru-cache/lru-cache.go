package leetcode

type LRUCache struct {
	hashMap map[int] *doubleLinkedList
	capacity int
	currentCount int
	head *doubleLinkedList
	tail *doubleLinkedList
}

type doubleLinkedList struct {
	value int
	key int
	pre *doubleLinkedList
	next *doubleLinkedList
}

func Constructor(capacity int) LRUCache {
	head := &doubleLinkedList{value: 0, key:   0,}
	tail := &doubleLinkedList{value: 0, key:   0,}

	head.next = tail
	tail.pre = head
	return LRUCache{
		hashMap: map[int]*doubleLinkedList{},
		capacity: capacity,
		currentCount: 0 ,
		head:     head,
		tail:     tail,
	}
}


func (this *LRUCache) Get(key int) int {

	if node, ok := this.hashMap[key] ; ok {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.next = this.head.next
		this.head.next.pre = node
		node.pre = this.head
		this.head.next = node
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if node , ok := this.hashMap[key] ; ok {
		node.value = value
		node.pre.next = node.next
		node.next.pre = node.pre
		node.next = this.head.next
		this.head.next.pre = node
		node.pre = this.head
		this.head.next = node
	} else {
		tmp := &doubleLinkedList{key:key, value:value}
		this.hashMap[key] = tmp
		tmp.pre = this.head
		tmp.next = this.head.next
		this.head.next.pre = tmp
		this.head.next = tmp
		this.currentCount ++
		if this.currentCount > this.capacity {
			delete(this.hashMap,this.tail.pre.key)
			this.tail.pre.pre.next = this.tail.pre.next
			this.tail.pre.next.pre = this.tail.pre.pre
			this.currentCount --
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
