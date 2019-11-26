package trie

import "fmt"

type TrieNode struct {
	children []*TrieNode

	data         byte
	isEndingChar bool
}

func NewTrie(data byte) (tn TrieNode) {
	tn = TrieNode{
		data:     data,
		children: make([]*TrieNode, 26),
	}

	return
}

func (tn *TrieNode) Insert(data string) {
	p := tn

	for i := 0; i < len(data); i++ {
		index := data[i] - 'a'

		if p.children[index] == nil {
			p.children[index] = &TrieNode{
				data: data[i],
			}
		}

		p = p.children[index]
	}

	p.isEndingChar = true
}

func (tn *TrieNode) Find(pattern string) bool {
	p := tn

	for i := 0; i < len(pattern); i++ {
		index := pattern[i] - 'a'

		if p.children[index] == nil {
			return false // 不存在pattern
		}

		p = p.children[index]
	}

	if p.isEndingChar == false { // 不能完全匹配，只是前缀
		return false
	}

	return true //找到pattern

}

func (tn *TrieNode) String() string {
	result := "head"

	for tn != nil && tn.children != nil {
		result += fmt.Sprintf("<-%+v", string(tn.data))
	}

	result += "<-tail"

	return result
}
