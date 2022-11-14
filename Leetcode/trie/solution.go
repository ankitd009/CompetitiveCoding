// https://leetcode.com/problems/implement-trie-prefix-tree/

package main

const (
	a = 97
)

type Node struct {
	word  bool
	nodes map[int]*Node
}

func NewNode() *Node {
	return &Node{
		nodes: make(map[int]*Node),
	}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	t := Trie{
		root: NewNode(),
	}
	return t
}

func (this *Trie) Insert(word string) {

	node := this.root

	for i, ch := range word {

		index := int(ch - a)

		if node.nodes[index] == nil {
			n := NewNode()
			node.nodes[index] = n
		}

		if len(word)-1 == i {
			n := node.nodes[index]
			n.word = true
			node.nodes[index] = n
		}

		node = node.nodes[index]
	}
}

func (this *Trie) Search(word string) bool {
	node := this.root

	for i := 0; i < len(word); i++ {

		ch := word[i]
		index := int(ch - a)

		if node.nodes[index] == nil {
			return false
		}

		node = node.nodes[index]
	}

	return node != nil && node.word
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	for i := 0; i < len(prefix); i++ {

		ch := prefix[i]
		index := int(ch - a)

		if node.nodes[index] == nil {
			return false
		}

		node = node.nodes[index]
	}

	return node != nil
}
