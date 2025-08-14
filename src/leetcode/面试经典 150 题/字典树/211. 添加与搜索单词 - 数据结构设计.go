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
		if node[index] == nil {
			node[index] = &letterNode{
				children: make([]*letterNode, 26),
			}
		}

		if i == len(word)-1 {
			node[index].isLeaf = true
		}
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

	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < len(node); j++ {
				if node[j] != nil && this.dfs(node[j].children, word[1:]) {
					return true
				}
			}
			return false
		}

		if node[int(word[i]-'a')] == nil {
			return false
		}

		return this.dfs(node[int(word[i]-'a')].children, word[1:])

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

// ["WordDictionary","addWord","addWord","addWord","addWord","search","search","addWord","search","search","search","search","search","search"]
// [[],["at"],["and"],["an"],["add"],["a"],[".at"],["bat"],[".at"],["an."],["a.d."],["b."],["a.d"],["."]]
// [null,null,null,null,null,false,false,null,true,false,false,true,false,true]
// [null,null,null,null,null,false,false,null,true,true,false,false,true,false]

// type TrieNode struct {
// 	children [26]*TrieNode
// 	isEnd    bool
// }

// func (t *TrieNode) Insert(word string) {
// 	node := t
// 	for _, ch := range word {
// 		ch -= 'a'
// 		if node.children[ch] == nil {
// 			node.children[ch] = &TrieNode{}
// 		}
// 		node = node.children[ch]
// 	}
// 	node.isEnd = true
// }

// type WordDictionary struct {
// 	trieRoot *TrieNode
// }

// func Constructor() WordDictionary {
// 	return WordDictionary{&TrieNode{}}
// }

// func (d *WordDictionary) AddWord(word string) {
// 	d.trieRoot.Insert(word)
// }

// func (d *WordDictionary) Search(word string) bool {
// 	var dfs func(int, *TrieNode) bool
// 	dfs = func(index int, node *TrieNode) bool {
// 		if index == len(word) {
// 			return node.isEnd
// 		}
// 		ch := word[index]
// 		if ch != '.' {
// 			child := node.children[ch-'a']
// 			if child != nil && dfs(index+1, child) {
// 				return true
// 			}
// 		} else {
// 			for i := range node.children {
// 				child := node.children[i]
// 				if child != nil && dfs(index+1, child) {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(0, d.trieRoot)
// }
