package trie

import "testing"

func TestSkipList(t *testing.T) {
	trie := NewTrie()

	trie.Insert("leo")
	t.Log(
		trie.StartsWith("l"),
		trie.StartsWith("le"),
		trie.Find("leo"),

		trie.StartsWith("lo"),
	)

	t.Log("-----------------------------")
	t.Log("未插入abc，查找abc", trie.Find("abc"))
	trie.Insert("abc")
	t.Log("插入abc，查找abc", trie.Find("abc"))

	trie.Insert("apple")
	t.Log(trie.Find("apple"))     // 返回 true
	t.Log(trie.Find("app"))       // 返回 false
	t.Log(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	t.Log(trie.Find("app")) // 返回 true

}
