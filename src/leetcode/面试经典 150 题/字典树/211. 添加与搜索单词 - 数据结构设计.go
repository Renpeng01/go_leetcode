package main

import "fmt"

type letterNode struct {
	isLeaf   bool
	children []*letterNode
}

type WordDictionary struct {
	root []*letterNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: make([]*letterNode, 26),
	}
}

func (this *WordDictionary) AddWord(word string) {

	node := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if word[i] == '.' {
			continue
		}

		if node[index] == nil {
			node[index] = &letterNode{
				isLeaf:   i == len(word)-1,
				children: make([]*letterNode, 26),
			}
		}

		if i == len(word)-1 {
			node[index].isLeaf = true
		}

		// fmt.Println(index, node[index].isLeaf)
		node = node[index].children
	}

}

func (this *WordDictionary) Search(word string) bool {
	node := this.root
	return this.dfs(node, word)
}

func (this *WordDictionary) dfs(node []*letterNode, word string) bool {
	if node == nil {

		return false
	}

	if len(word) == 1 && (word[0] == '.' || node[int(word[0]-'a')] != nil && node[int(word[0]-'a')].isLeaf) {

		return true
	}

	res := false
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < 26; j++ {
				if node[j] != nil {
					res = res || this.dfs(node[j].children, word[1:])
				}
			}
			return res
		}

		if node[int(word[i]-'a')] == nil {

			return false
		}
		if i == len(word)-1 && node[int(word[i]-'a')].isLeaf {
			return true
		}
		node = node[int(word[i]-'a')].children
	}

	return res
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()
	// wordDictionary.AddWord("bad")
	// wordDictionary.AddWord("dad")
	// wordDictionary.AddWord("mad")

	// res := wordDictionary.Search("pad") // 返回 False
	// fmt.Println(res)

	// res = wordDictionary.Search("bad") // 返回 True
	// fmt.Println(res)

	// res = wordDictionary.Search(".ad") // 返回 True
	// fmt.Println(res)

	// res = wordDictionary.Search("b..") // 返回 True
	// fmt.Println(res)

	wordDictionary.AddWord("a")
	wordDictionary.AddWord("ab")

	res := wordDictionary.Search("a")
	fmt.Println(res)
}
