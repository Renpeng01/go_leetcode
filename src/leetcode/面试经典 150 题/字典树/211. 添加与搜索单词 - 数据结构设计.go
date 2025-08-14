package main

import "fmt"

type letterNode struct {
	c        byte
	isLeaf   bool
	children []*letterNode
}

type WordDictionary struct {
	root []*letterNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: make([]*letterNode, 27),
	}
}

func (this *WordDictionary) AddWord(word string) {

	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if word[i] == '.' {
			index = 26
		}

		node[index] = &letterNode{
			c:        word[i],
			isLeaf:   i == len(word)-1,
			children: make([]*letterNode, 27),
		}
		node = node[index].children
	}

}

func (this *WordDictionary) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if word[i] == '.' {
			index = 26
		}

		n := node[index]
		if n == nil {
			return false
		}

		if i == len(word)-1 && n.isLeaf {
			return true
		}

		node = n.children
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	res := wordDictionary.Search("pad") // 返回 False
	fmt.Println(res)

	res = wordDictionary.Search("bad") // 返回 True
	fmt.Println(res)

	res = wordDictionary.Search(".ad") // 返回 True
	fmt.Println(res)

	res = wordDictionary.Search("b..") // 返回 True
	fmt.Println(res)
}
