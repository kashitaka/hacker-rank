package _55_min_stack

type Node struct {
	Val int
	Min int
}

type MinStack struct {
	stack []*Node
}

func Constructor() MinStack {
	stack := make([]*Node, 0)
	return MinStack{
		stack: stack,
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.stack) > 0 {
		curMin := this.stack[len(this.stack)-1].Min
		if curMin < val {
			min = curMin
		}
	}
	new := &Node{Val: val, Min: min}
	this.stack = append(this.stack, new)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
