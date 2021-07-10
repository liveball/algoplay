/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
    var freq [256]int
	res:=false

	if len(s2)==0 || len(s2)<len(s1){
		return res
	}

	//统计s1每个字母出现频次
	for i:=0;i<len(s1);i++{
		freq[s1[i]-'a']++
	}

	left,right,count:=0,0,len(s1)

	for right<len(s2){
		//s2中第一个出现的字母，p总数减1
		if freq[s2[right]-'a']>=1{
			count--
		}
		freq[s2[right]-'a']--//出现频次减1
		right++//继续右边滑动

		if count==0{//左边界移出不符合规范的元素
            res=true
		}

		//判断每个字母是否都被访问过一遍了
		if right-left==len(s1){
			if freq[s2[left]-'a']>=0{
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}

	return res
}
// @lc code=end

