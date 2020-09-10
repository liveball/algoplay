package main

import "fmt"

type MaxStack struct {
	elements, max []int
}

/** initialize your data structure here. */
func ConstructorMaxStack() MaxStack {
	return MaxStack{make([]int, 0), make([]int, 0)}
}

func (this *MaxStack) Push(x int) {
	this.elements = append(this.elements, x)

	if len(this.max) == 0 {
		this.max = append(this.max, x)
	} else {
		tmax := this.PeekMax()
		if tmax < x {
			this.max = append(this.max, x)
		} else {
			this.max = append(this.max, tmax)
		}
	}

}

func (this *MaxStack) Pop() int {
	l := len(this.elements)
	if l == 0 {
		return 0
	}

	elem := this.elements[l-1]
	this.elements = this.elements[:l-1]
	return elem
}

func (this *MaxStack) Top() int {
	l := len(this.elements)
	if l == 0 {
		return 0
	}
	return this.elements[l-1]
}

func (this *MaxStack) PeekMax() int {
	l := len(this.max)
	if l == 0 {
		return 0
	}
	return this.max[l-1]
}

func (this *MaxStack) PopMax() int {
	l := len(this.max)
	if l == 0 {
		return 0
	}

	tmax := this.max[l-1]
	this.max = this.max[:l-1]
	return tmax
}

func (this *MaxStack) PrintMaxStack() {
	fmt.Println(this.elements, this.max)
}
