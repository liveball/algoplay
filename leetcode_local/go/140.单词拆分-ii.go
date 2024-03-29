/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}

// @lc code=end

