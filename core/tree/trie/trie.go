package trie

import "strings"

type TrieNode struct {
	children []*TrieNode

	isEndingChar bool
}

func NewTrie() (tn *TrieNode) {
	tn = &TrieNode{
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
				children: make([]*TrieNode, 26),
			}
		}

		p = p.children[index]
	}

	p.isEndingChar = true
	strings.ToLower(data)
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

func (tn *TrieNode) StartsWith(prefix string) bool {
	p := tn

	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'

		if ch := p.children[index]; ch == nil {
			return false
		} else {
			p = ch
		}
	}

	return true
}

//func (tn *TrieNode) String() string {
//	result := "head"
//
//	for tn != nil && len(tn.children)>0 {
//		for key := range tn.children {
//			result += fmt.Sprintf("<-%+v", string(key))
//		}
//	}
//
//	result += "<-tail"
//
//	return result
//}
