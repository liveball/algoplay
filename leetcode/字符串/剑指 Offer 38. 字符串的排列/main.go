package main

func main(){}


func permutation(s string) []string {
    n := len(s)
    visited := make([]bool, n)
    res := make([]string, 0)
    cur := make([]byte, n)
    var dfs func(idx int)
    dfs = func(idx int) {
        if idx == n {
            res = append(res, string(cur))
            return
        }

        d := make(map[byte]bool)
        for i, _ := range s {
            if _, ok := d[s[i]]; ok {
                continue
            }

            if visited[i] {
                continue
            }

            visited[i] = true
            d[s[i]] = true
            cur[idx] = s[i]
            dfs(idx + 1)
            visited[i] = false
        }
    }

    dfs(0)
    return res
}
// 输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
// 示例:

// 输入：s = "abc"
// 输出：["abc","acb","bac","bca","cab","cba"]

// 限制：

// 1 <= s 的长度 <= 8
