package main

// type TrieNoe struct {
// 	isLeaf bool
// 	vals   []*TrieNoe
// }

// type Trie struct {
// 	root *TrieNoe
// }

// func Constructor() Trie {
// 	tire := Trie{
// 		root: &TrieNoe{},
// 	}
// 	return tire
// }

// func (this *Trie) Insert(word string) {
// 	node := this.root
// 	for i := 0; i < len(word); i++ {
// 		if node.vals == nil {
// 			node.vals = make([]*TrieNoe, 26)
// 		}

// 		if node.vals[word[i]-97] == nil {
// 			node.vals[word[i]-97] = &TrieNoe{} // 赋值
// 		}
// 		if i == len(word)-1 {
// 			node.isLeaf = true
// 		}
// 		node = node.vals[word[i]-97]
// 	}
// }

// func (this *Trie) Search(word string) bool {
// 	node := this.find(word)
// 	return node != nil && node.isLeaf

// }

// func (this *Trie) find(word string) *TrieNoe {
// 	node := this.root
// 	for i := 0; i < len(word); i++ {
// 		if len(node.vals) == 0 || node.vals[word[i]-97] == nil {
// 			return nil
// 		}
// 		if i == len(word)-1 {
// 			break
// 		}
// 		node = node.vals[word[i]-97]
// 	}
// 	return node
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	return this.find(prefix) != nil
// }

// func main() {
// 	aa := 'a'
// 	fmt.Println(aa) // 97
// }

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
