package main

type MinStack struct {
	st    []int // 正常栈
	minSt []int // 最小栈：栈顶为当前最小值
}

func Constructor() MinStack {
	return MinStack{
		st:    make([]int, 0),
		minSt: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)

	// minSt 维护“到当前为止”的最小值
	if len(this.minSt) == 0 || val <= this.minSt[len(this.minSt)-1] {
		this.minSt = append(this.minSt, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.st) == 0 {
		return // 按题意通常不会对空栈调用；这里防御式处理
	}

	top := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]

	if len(this.minSt) > 0 && top == this.minSt[len(this.minSt)-1] {
		this.minSt = this.minSt[:len(this.minSt)-1]
	}
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}
