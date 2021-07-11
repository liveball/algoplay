/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	min     []int
	element []int
	size    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:     make([]int, 0),
		element: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.element = append(this.element, val)

	if this.size == 0 {
		this.min = append(this.min, val)
	} else {
		x := this.GetMin()
		if x < val {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, val)
		}
	}

	this.size++
}

func (this *MinStack) Pop() {
	this.size--
	this.min = this.min[:this.size]
	this.element = this.element[:this.size]
}

func (this *MinStack) Top() int {
	return this.element[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

