package _46_lru_cache

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Cap   int
	Hash  map[int]*Node
	Left  *Node
	Right *Node
}

func Constructor(capacity int) LRUCache {
	hash := make(map[int]*Node)
	left, right := &Node{}, &Node{}
	left.Next = right
	right.Prev = left
	return LRUCache{
		Cap:   capacity,
		Hash:  hash,
		Left:  left,
		Right: right,
	}
}

func (c *LRUCache) insert(node *Node) {
	prev := c.Right.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = c.Right
	c.Right.Prev = node
}

func (c *LRUCache) remove(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.Hash[key]
	if !ok {
		return -1
	}
	c.remove(node)
	c.insert(node)
	return node.Val
}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.Hash[key]
	if ok {
		c.remove(node)
	}
	node = &Node{Key: key, Val: value}
	c.insert(node)
	c.Hash[key] = node

	if len(c.Hash) > c.Cap {
		lru := c.Left.Next
		c.remove(lru)
		delete(c.Hash, lru.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
