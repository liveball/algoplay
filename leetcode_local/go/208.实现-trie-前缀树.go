/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children []*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	ts := this
	for _, w := range word {
		if ts.children[w-'a'] == nil {
			tr := Constructor()
			ts.children[w-'a'] = &tr
		}

		ts = ts.children[w-'a']
	}

	ts.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ts := this
	for _, w := range word {
		if ts.children[w-'a'] == nil {
			return false
		}

		ts = ts.children[w-'a']
	}

	if ts.isEnd == false { //不能完全匹配， 只是一个前缀
		return false
	}

	return ts.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ts := this
	for _, w := range prefix {
		if ch := ts.children[w-'a']; ch == nil {
			return false
		} else {
			ts = ch
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

