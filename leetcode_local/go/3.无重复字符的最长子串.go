/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	return slidingWindow(s)
}

func slidingWindow(s string) int {
	if len(s)==0{
		return 0
	}
    
	res := 0
	left, right := 0, 0
    var bitSet [256]uint8
	for left< len(s){
		//没有重复的字符，扩大右边界
        if right<len(s) && bitSet[s[right]]==0{
           bitSet[s[right]] = 1 //accii 标记已访问的边界
		   right++
		} else {//出现重复的字符
			bitSet[s[left]] = 0//标记未访问的边界
			left++ //缩小左边界，将重复字符移出左边界
		}

		res = max(res, right-left)
	}


	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

