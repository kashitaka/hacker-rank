package main

type DoubleLinkedNode struct {
	key, value int
	prev, next *DoubleLinkedNode
}

type LRUCache struct {
	cache          map[int]*DoubleLinkedNode
	size, capacity int
	head, tail     *DoubleLinkedNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	cache := make(map[int]*DoubleLinkedNode, capacity)
	lruCache.capacity = capacity
	lruCache.cache = cache
	lruCache.head = &DoubleLinkedNode{}
	lruCache.tail = &DoubleLinkedNode{}
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (this *LRUCache) addNode(node *DoubleLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoubleLinkedNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *DoubleLinkedNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *DoubleLinkedNode {
	res := this.tail.prev
	this.removeNode(res)
	return res
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		// そのまま上書きし、アクセスを更新
		node.value = value
		this.moveToHead(node)
		return
	}
	// create new k-v
	newNode := DoubleLinkedNode{
		key:   key,
		value: value,
	}
	if this.size < this.capacity {
		// そのまま突っ込む
		this.addNode(&newNode)
		this.size++
		this.cache[key] = &newNode
		return
	}
	// 一番古いのを取り出して、新しいのを突っ込む
	tail := this.popTail()
	delete(this.cache, tail.key)
	this.addNode(&newNode)
	this.cache[key] = &newNode
}
