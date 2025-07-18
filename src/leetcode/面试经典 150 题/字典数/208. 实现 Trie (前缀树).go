package main

import (
	"fmt"
)

type TrieNoe struct {
	childern []*TrieNoe
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
		if len(node.childern) == 0 {
			node.childern = make([]*TrieNoe, 26)
		}

		if node.childern[word[i]-97] == nil {
			node.childern[word[i]-97] = &TrieNoe{}
		}

		node = node.childern[word[i]-97]
	}
}

func (this *Trie) Search(word string) bool {
	node := this.find(word)
	return node != nil && len(node.childern) == 0

}

func (this *Trie) find(word string) *TrieNoe {
	node := this.root
	for i := 0; i < len(word); i++ {
		if len(node.childern) == 0 || node.childern[word[i]-97] == nil {
			return nil
		}
		node = node.childern[word[i]-97]
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
