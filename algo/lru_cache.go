package main

import "fmt"

func main() {
	fmt.Printf("%+v", "Hello World")
}

func NewList()*List{
	return &List{
		CurSize: 0,
		Head:    nil,
		Tail:    nil,
	}
}

type List struct {
	CurSize int
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val string
	Key string
	Next *ListNode
	Prev *ListNode
}

func (l *List) insertKVToEnd(key, value string)*ListNode {
	node := &ListNode{
		Val:  value,
		Key:  key,
		Next: nil,
		Prev: nil,
	}
	return l.insertNodeToEnd(node)
}

func (l *List) insertNodeToEnd(node *ListNode)*ListNode {

	node.Prev = l.Tail

	if l.Tail != nil {
		l.Tail.Next = node
	}

	l.Tail = node

	// do we update head ?
	if l.CurSize == 0 {
		l.Head = node
	}

	l.CurSize = l.CurSize + 1

	return node
}

func (l *List) deleteFromStart()*ListNode {
	node := l.Head
	l.Head = l.Head.Next
	if l.Head.Next == nil {
		l.Tail = nil
	}else {
		l.Head.Next.Prev = nil
	}
	l.CurSize = l.CurSize - 1
	node.Prev = nil
	node.Next = nil
	return node
}

func (l *List) deleteNode(node *ListNode) {
	if l.CurSize == 0 {
		return
	}
	prev := node.Prev
	next := node.Next

	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}

	node.Prev = nil
	node.Next = nil
	l.CurSize = l.CurSize - 1
}

type LRUCache struct {
	cacheList *List
	keyToListNode map[string]*ListNode
	size int
}

func NewLRUCache(size int)*LRUCache {
	return &LRUCache{
		cacheList:     NewList(),
		keyToListNode: make(map[string]*ListNode),
		size:          size,
	}
}

func (lru *LRUCache) Insert(k ,v string){
	node, exists := lru.keyToListNode[k]
	if exists {
		node.Val = v
		lru.cacheList.deleteNode(node)
		lru.cacheList.insertNodeToEnd(node)
		return
	}

	// when size full
	if lru.cacheList.CurSize == lru.size {
		deletedNode := lru.cacheList.deleteFromStart()
		delete(lru.keyToListNode, deletedNode.Key)
	}

	node = lru.cacheList.insertKVToEnd(k,v)
	lru.keyToListNode[k] = node
	lru.size = lru.size + 1
}

func (lru *LRUCache) Get(k string)string{
	// cache miss
	node, exists := lru.keyToListNode[k]
	if !exists {
		return ""
	}

	// refresh usage stats
	lru.cacheList.deleteNode(node)
	lru.cacheList.insertNodeToEnd(node)

	return node.Val
}

