package main

type TrieNode struct {
	data         byte
	children     [26]*TrieNode
	isEndingChar bool
}

type Trie struct {
	root *TrieNode
}

func (trie *Trie) insert(text []byte) {
	p := trie.root
	for i := 0; i < len(text); i++ {
		index := text[i] - 'a'
		if p.children[index] == nil {
			p.children[index] = &TrieNode{
				data: text[i],
			}
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func (trie *Trie) find(pattern []byte) bool {
	p := trie.root
	for i := 0; i < len(pattern); i++ {
		index := pattern[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	return p.isEndingChar
}
