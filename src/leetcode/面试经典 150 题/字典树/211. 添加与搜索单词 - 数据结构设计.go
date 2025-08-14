package main

type letterNode struct {
	c        byte
	isLeaf   bool
	children [26]*letterNode
}

type WordDictionary struct {
	root [26]*letterNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: [26]*letterNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {

	node := this.root
	for i := 0; i < len(word); i++ {
		node[word[i]-'a'] = &letterNode{
			c:        word[i],
			isLeaf:   i == len(word)-1,
			children: [26]*letterNode{},
		}
		node = node[word[i]-'a'].children
	}

}

func (this *WordDictionary) Search(word string) bool {

	node := this.root
	for i := 0; i < len(word); i++ {
		n := node[int(word[i]-'a')]
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
