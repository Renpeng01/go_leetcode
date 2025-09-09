package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	queue := MyQueue{
		stack1: make([]int, 0, 32),
		stack2: make([]int, 0, 32),
	}
	return queue
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {

	// v := this.stack1[len(this.stack1)-1]
	// this.stack1 = this.stack1[:len(this.stack1)-1]
	// return
}

func (this *MyQueue) Peek() int {

}

func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
