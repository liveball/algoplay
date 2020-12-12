package main

func main(){}

func isPalindrome(s string) bool {
    // 空是有效的回文
	if s == "" {
		return true
	}
	// 统一转换为小写
	s = strings.ToLower(s)
	// 利用strings builder来构建string
	var build strings.Builder
	// 只要 a-z 和 0-9的字符
	for i := range s {
		if ('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9') {
			build.WriteByte(s[i])
		}
	}
	// 得到结果
	s = build.String()

    i:=0
    j:=len(s)-1

    for i<j {
        if s[i]!=s[j]{
            return false
        }
        i++
        j--
    }

    return true
}

func palindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-palindrome
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。