package trie

import "testing"

func TestSkipList(t *testing.T) {
	trie := NewTrie('/')

	trie.Insert("leo")
	t.Log(trie)
	t.Log("-----------------------------")

}
