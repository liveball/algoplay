/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	Stack *[]int
	Queue *[]int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	tmp1, tmp2 := []int{}, []int{}
	return MyQueue{
		Stack: &tmp1,
		Queue: &tmp2,
	}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.Stack = append(*this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*this.Queue)==0{
       this.stack2Queue(this.Stack, this.Queue)
	}

	poped:=(*this.Queue)[len(*this.Queue)-1]
	*this.Queue=(*this.Queue)[:len(*this.Queue)-1]
	return poped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.Queue)==0{
		this.stack2Queue(this.Stack, this.Queue)
	 }

	return (*this.Queue)[len(*this.Queue)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.Stack)+len(*this.Queue) == 0
}

func (this *MyQueue) stack2Queue(stack, queue *[]int) {
	for len(*stack) > 0 {
		poped := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		*queue = append(*queue, poped)
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

