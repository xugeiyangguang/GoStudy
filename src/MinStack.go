package main

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) != 0 && this.minStack[len(this.minStack)-1] < x {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	} else {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	r1 := obj.GetMin()
	println(r1)
	obj.Pop()
	r3 := obj.Top()
	println(r3)
	r4 := obj.GetMin()
	println(r4)

}
