package unitTestExcercise

type LRUCache struct {
	head *cache_node
	last *cache_node
	cap  int
	mapp map[int]*cache_node
}

func Constructor(capacity int) LRUCache {
	head := &cache_node{}
	last := &cache_node{}
	head.next = last
	last.pre = head
	mapp := make(map[int]*cache_node)
	cache := LRUCache{head: head, last: last, cap: capacity, mapp: mapp}
	return cache
}

func (this *LRUCache) Get(key int) int {
	tmp, ok := this.mapp[key]
	if !ok {
		return -1
	}
	tmp.next.pre = tmp.pre
	tmp.pre.next = tmp.next
	tmp.next = this.head.next
	tmp.pre = this.head
	this.head.next.pre = tmp
	this.head.next = tmp
	return tmp.val
}

func (this *LRUCache) Put(key int, value int) {
	tmp, ok := this.mapp[key]
	if ok {
		tmp.val = value
		tmp.next.pre = tmp.pre
		tmp.pre.next = tmp.next
		tmp.next = this.head.next
		tmp.pre = this.head
		this.head.next.pre = tmp
		this.head.next = tmp
	} else {
		newNode := &cache_node{val: value, key: key, pre: this.head, next: this.head.next}
		this.head.next.pre = newNode
		this.head.next = newNode
		this.mapp[key] = newNode
		if this.cap < len(this.mapp) { //超出容量
			del := this.last.pre.key
			this.last.pre.pre.next = this.last
			this.last.pre = this.last.pre.pre
			delete(this.mapp, del)
		}
	}

}

type cache_node struct {
	key  int
	val  int
	pre  *cache_node
	next *cache_node
}
