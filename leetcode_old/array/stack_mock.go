package main

type MyStack struct {
	a []int
}


/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	if len(this.a)==0{
		this.a=make([]int,0)
	}

	this.a=append(this.a, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.a)==0{
		return 0
	}

	t:=this.a[len(this.a)-1]
	this.a=this.a[:len(this.a)-1]

	return t
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.a[len(this.a)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.a)==0{
		return true
	}

	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
