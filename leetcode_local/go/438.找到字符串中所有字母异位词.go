/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
    var freq [256]int
	res:=[]int{}

	if len(s)==0 || len(s)<len(p){
		return res
	}

	//统计p每个字母出现频次
	for i:=0;i<len(p);i++{
		freq[p[i]-'a']++
	}

	left,right,count:=0,0,len(p)

	for right<len(s){
		//s中第一个出现的字母，p总数减1
		if freq[s[right]-'a']>=1{
			count--
		}
		freq[s[right]-'a']--//出现频次减1
		right++//继续右边滑动

		if count==0{//左边界移出不符合规范的元素
            res=append(res, left)
		}

		//判断每个字母是否都被访问过一遍了
		if right-left==len(p){
			if freq[s[left]-'a']>=0{
				count++
			}
			freq[s[left]-'a']++
			left++
		}

	}

	return res

}
// @lc code=end

