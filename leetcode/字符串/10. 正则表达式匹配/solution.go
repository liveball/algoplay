package _10_正则表达式匹配

//动态规划三部
//设dp[i][j]表示p[:i]能否正则匹配s[:j]，m为p的长度，n为s的长度，则dp[m][n]即为所求
//
//状态转移方程
//p[i-1] == s[j-1] || p[i-1] == '.'时, dp[i][j] = dp[i-1][j-1]
//dp[i-1]=='*'时，分两种情况
//a. p[i-2] == s[j-1] || p[i-2] == '.'时, dp[i][j] = dp[i][j-1] || dp[i-2][j-1] || dp[i-2][j]
//b. 否则，dp[i][j] = dp[i-2][j]
//初始值
//可以确定dp[i][0], dp[0][j]得值
//
//求解dp[0][j]
//
//当j=0时，dp[0][0]=true
//否则dp[0][j]=false
//求解dp[i][0]
//
//当i为基数时，dp[j][0]=false
//当i为偶数时并且p[i-1] == '*'时，dp[i][0]=dp[i-2][0]

func isMatch(s string, p string) bool {
	m, n := len(p)+1, len(s)+1
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}

	//初始值
	dp[0][0] = true
	for i := 2; i < m; i++ {
		if i%2 == 0 && p[i-1] == '*' {
			dp[i][0] = dp[i-2][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[i-1] == s[j-1] || p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' {
				if i > 1 { //防止脚标越界
					if p[i-2] == s[j-1] || p[i-2] == '.' {
						dp[i][j] = dp[i][j-1] || dp[i-2][j-1] || dp[i-2][j]
					} else {
						dp[i][j] = dp[i-2][j]
					}
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func isMatch2(s string, p string) bool {
	// 新建二维状态数组作为备忘录，存储 s[i] p[j] 之后的匹配情况
	// m[i][j] 代表 s[i] 和 p[j] 之后是否匹配
	// 1 代表匹配， -1 代表不匹配
	var m [][]int
	// 初始化状态数组
	for i := 0; i <= len(s); i++ {
		tmp := make([]int, len(p)+1)
		m = append(m, tmp)
	}
	return match(0, 0, s, p, m)
}

func match(i, j int, s, p string, m [][]int) bool {
	// 如果备忘录中已有记录，则直接返回
	if m[i][j] != 0 {
		return m[i][j] == 1 // 1 代表匹配
	}
	var result bool
	if j >= len(p) {
		result = i >= len(s)
	} else {
		curMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
		// 如果 j 后面是 *
		if j+1 < len(p) && p[j+1] == '*' {
			result = match(i, j+2, s, p, m) || (curMatch && match(i+1, j, s, p, m))
			// j 后面不是 * 或者 j+1 >= len(p)
		} else {
			// 当前匹配，后面的也匹配
			result = curMatch && match(i+1, j+1, s, p, m)
		}
	}
	// 更新备忘录
	if result {
		m[i][j] = 1
	} else {
		m[i][j] = -1
	}
	return result
}
