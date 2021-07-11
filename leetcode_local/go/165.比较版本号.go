/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	l := len(v1)
	for i := 0; i < l; i++ {
		vs1 := strings.TrimLeft(v1[i], "0")
		vs2 := strings.TrimLeft(v2[i], "0")

		for len(vs1) < len(vs2) {
			vs1 = "0" + vs1
		}
		for len(vs2) < len(vs1) {
			vs2 = "0" + vs2
		}

		vl := len(vs1)
		for j := 0; j < vl; j++ {
			if vs1[j] == vs2[j] {
				continue
			}

			if vs1[j] < vs2[j] {
				return -1
			}

			return 1
		}
	}

	return 0
}

// @lc code=end

