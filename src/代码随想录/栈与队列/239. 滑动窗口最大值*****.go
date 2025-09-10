package main

import (
	"fmt"
)

type Queue struct {
	q []int
}

func (this *Queue) push(v int) {
	for len(this.q) > 0 && this.q[len(this.q)-1] < v {
		this.q = this.q[:len(this.q)-1]
	}
	this.q = append(this.q, v)

}

func (this *Queue) pop(v int) {
	if len(this.q) > 0 && v == this.q[0] {
		this.q = this.q[1:]
	}
}

func (this *Queue) getMax() int {
	return this.q[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := &Queue{
		q: make([]int, 0, 16),
	}
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}

	res := make([]int, 0, 16)
	res = append(res, queue.getMax())

	fmt.Println("start q: ", queue.q)

	l, r := 0, k
	for r < len(nums) {
		queue.push(nums[r])
		fmt.Println("after push nums[r]:", nums[r], " q:", queue.q)
		queue.pop(nums[l])
		fmt.Println("after pop nums[r]:", nums[l], " q:", queue.q)

		res = append(res, queue.getMax())
		l++
		r++
	}

	return res
}

func main() {
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k := 3

	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3

	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
