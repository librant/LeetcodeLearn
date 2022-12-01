package main

import "fmt"

// https://leetcode.cn/problems/lru-cache/?favorite=2cktkvj

// LRU：双向链表 + Hash 表

func main() {
	lc := Constructor(2)
	lc.Put(1, 1)
	fmt.Printf("%+v\n", lc)
	lc.Put(2, 2)
	fmt.Printf("%+v\n", lc)
	v1 := lc.Get(1)
	fmt.Printf("v1:%v %+v\n", v1, lc)
	lc.Put(3, 3)
	fmt.Printf("%+v\n", lc)
	v2 := lc.Get(2)
	fmt.Printf("v2:%v %+v\n", v2, lc)
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	cache      map[int]*DLinkedNode // 缓存 key/value
	head, tail *DLinkedNode         // 数组的最前位置表示最少访问，数组最后的位置越最近访问
	size       int
	capacity   int // 当前缓存的容量
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	// 将当前访问过的节点移动到头节点
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	// 先将节点从链表中删除
	this.removeNode(node)
	// 再讲节点增加到链表头
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
