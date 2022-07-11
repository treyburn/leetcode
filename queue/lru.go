package queue

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type doublyLinkedList struct {
	key  int
	val  int
	prev *doublyLinkedList
	next *doublyLinkedList
}

type LRUCache struct {
	set  map[int]*doublyLinkedList
	cap  int
	size int
	head *doublyLinkedList
	tail *doublyLinkedList
}

func Constructor(capacity int) LRUCache {
	s := make(map[int]*doublyLinkedList, capacity)
	head := doublyLinkedList{}
	tail := doublyLinkedList{}
	head.next = &tail
	tail.prev = &head

	return LRUCache{
		set:  s,
		cap:  capacity,
		size: 0,
		head: &head,
		tail: &tail,
	}
}

func (this *LRUCache) addNode(node *doublyLinkedList) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *doublyLinkedList) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *doublyLinkedList) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *doublyLinkedList {
	last := this.tail.prev
	this.removeNode(last)
	return last
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.set[key]
	if !ok {
		return -1
	}
	this.moveToHead(v)

	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	dll, ok := this.set[key]
	if !ok {
		dll = &doublyLinkedList{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		this.addNode(dll)
		this.size++
		this.set[key] = dll
	} else {
		dll.val = value
		this.moveToHead(dll)
	}

	if this.size > this.cap {
		last := this.popTail()
		delete(this.set, last.key)
		this.size--
	}
}
