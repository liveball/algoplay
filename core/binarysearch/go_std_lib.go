package main

func SearchByStd(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)

		if !f(h) {
			i = h + 1 //在后半段查找
		} else {
			j = h //不退出循环，在前半段继续查找
		}
	}
	return i
}
