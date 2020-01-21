package main

func main() {

}

func climb(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 { //1个台阶
		return 1
	}

	if n == 2 { //两个台阶
		return 2
	}

	return climb(n-1) + climb(n-2) //n-2个台阶
}

func climbMemo(n int, tmap map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 { //1个台阶
		return 1
	}

	if n == 2 { //两个台阶
		return 2
	}

	if v, ok := tmap[n]; ok {
		return v
	} else {
		v = climbMemo(n-1, tmap) + climbMemo(n-2, tmap)
		tmap[n] = v
		return v
	}
}

//动态规划算法
func climbByDynamic(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 { //1个台阶
		return 1
	}

	if n == 2 { //两个台阶
		return 2
	}

	t, a, b := 0, 1, 2

	for i := 3; i <= n; i++ {
		t = a + b
		a = b
		b = t
	}

	return t
}
