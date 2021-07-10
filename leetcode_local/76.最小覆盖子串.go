/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
    if s=="" || t==""{
		return ""
	}

	var sFreq,tFreq [256]int
	res:=""
	left:=0
	right:=-1
	finalLeft:=-1
	finalRight:=-1
	minW:=len(s)+1
	count:=0

	for i:=0;i<len(t);i++{
		tFreq[t[i]-'a']++
	}

	for left < len(s){
		if right+1<len(s) && count<len(t){
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a']<=tFreq[s[right+1]-'a']{
				count++
			}
			right++
		} else {
			if right-left+1<minW  && count==len(t){
				minW=right-left+1
				finalLeft=left
				finalRight=right
			}

			if sFreq[s[left]-'a']==tFreq[s[left]-'a']{
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}

	if finalLeft!=-1 {
		for i:=finalLeft;i<finalRight+1;i++{
			res+=string(s[i])
		}
	}

	return res
}
// @lc code=end

