package main

import (
	"fmt"
)

type TrieNoe struct {
	isLeaf bool
	vals   []*TrieNoe
}

type Trie struct {
	root *TrieNoe
}

func Constructor() Trie {
	tire := Trie{
		root: &TrieNoe{},
	}
	return tire
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		if node.vals == nil {
			node.vals = make([]*TrieNoe, 26)
		}

		if node.vals[word[i]-97] == nil {
			node.vals[word[i]-97] = &TrieNoe{} // 赋值
		}
		if i == len(word)-1 {
			node.isLeaf = true
		}
		node = node.vals[word[i]-97]
	}
}

func (this *Trie) Search(word string) bool {
	node := this.find(word)
	return node != nil && node.isLeaf

}

func (this *Trie) find(word string) *TrieNoe {
	node := this.root
	for i := 0; i < len(word); i++ {
		if len(node.vals) == 0 || node.vals[word[i]-97] == nil {
			return nil
		}
		if i < len(word)-1 {
			node = node.vals[word[i]-97]
		}

	}
	return node
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}

func main() {
	aa := 'a'
	fmt.Println(aa) // 97
}
