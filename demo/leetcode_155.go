package main

// MinStack 是优化后的版本：辅助栈按需入栈（最终推荐版）
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor_155() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 只有当 val <= 当前最小值时才压入辅助栈
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	// 如果弹出的刚好是最小值，辅助栈也要同步弹出
	if val == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// MinStackBasic 展示了同步辅助栈方案（基础版）
type MinStackBasic struct {
	stack    []int
	minStack []int
}

func (this *MinStackBasic) Push(val int) {
	this.stack = append(this.stack, val)
	minVal := val
	if len(this.minStack) > 0 {
		if topMin := this.minStack[len(this.minStack)-1]; topMin < minVal {
			minVal = topMin
		}
	}
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStackBasic) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}
