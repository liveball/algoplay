/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// return largestRectangleArea1(heights)
	// return largestRectangleArea2(heights)
	return largestRectangleArea3(heights)
}

func largestRectangleArea1(heights []int) int {
	area := 0
	size := len(heights)
	for i := 0; i < size; i++ {
		left := i
		right := i

		for ; left >= 0; left-- {
			if heights[left] < heights[i] {
				break
			}
		}

		for ; right < size; right++ {
			if heights[right] < heights[i] {
				break
			}
		}

		w := right - left - 1
		h := heights[i]
		area = max(area, w*h)
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func largestRectangleArea2(heights []int) int {
	area := 0
	size := len(heights)
	minLeftIndex := make([]int, size)
	minRightIndex := make([]int, size)

	minLeftIndex[0] = -1
	for i := 1; i < size; i++ {
		t := i - 1
		for t >= 0 && heights[t] >= heights[i] {
			t = minLeftIndex[i]
		}
		minLeftIndex[i] = t
	}

	minRightIndex[size-1] = size
	for i := size - 2; i >= 0; i-- {
		t := i + 1
		for t < size && heights[t] >= heights[i] {
			t = minRightIndex[i]
		}
		minRightIndex[i] = t
	}

	for i := 0; i < size; i++ {
		area = max(area, heights[i]*(minRightIndex[i]-minLeftIndex[i]-1))
	}

	return area
}

func largestRectangleArea3(heights []int) int {
	area := 0
	//slice 头部加入元素0 方法1
	// heights = append([]int{0}, heights...)

	//slice 头部加入元素0  方法2
	heights = append(heights, 0)             //先把原来的切片长度+1
	index := 0                               //要把新元素插入到第二个位置
	copy(heights[index+1:], heights[index:]) // 从index位置开始拷贝
	heights[index] = 0                       //新元素的值是0

	// slice 尾部加入元素0
	heights = append(heights, 0)

	size := len(heights)
	stack := make([]int, size)
	for i := 0; i < size; i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				right := i
				left := stack[len(stack)-1]
				w := right - left - 1
				area = max(area, w*heights[mid])
			}
			stack = append(stack, i)
		}
	}

	return area

}

// @lc code=end

