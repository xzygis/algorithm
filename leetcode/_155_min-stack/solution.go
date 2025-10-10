package _155_min_stack

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val)
		this.min = append(this.min, val)
		return
	}

	this.stack = append(this.stack, val)
	if val < this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.GetMin())
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
