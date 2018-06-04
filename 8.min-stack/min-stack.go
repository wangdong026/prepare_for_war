package main

func main() {
	//mStack := MinStack{}
	mStack := new(MinStack)
	mStack.Push(-2)
	mStack.Push(0)
	mStack.Push(-3)
	println(mStack.GetMin())
	mStack.Pop()
	println(mStack.Top())
	println(mStack.GetMin())
}

type MinStack struct {
	stack []item
}

type item struct {
	Val int
	Min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{Val: x, Min: min})
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1].Val
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
