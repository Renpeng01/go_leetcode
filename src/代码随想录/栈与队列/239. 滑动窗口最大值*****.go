package main

type Queue struct {
	q []int
}

func (this *Queue) push(v int) {
	index := -1
	for i, m := range this.q {
		if m < v {
			index = i
		}
	}

	this.q = append(this.q, v)
	this.q = this.q[index+1:]
}

func (this *Queue) pop(v int) {
	if v == this.q[0] {
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
	res := make([]int, 0, len(nums))
	for i, v := range nums {
		if i < k-1 {
			queue.push(v)
			continue
		}
		queue.push(v)
		res = append(res, queue.getMax())
		queue.pop(v)
	}
	return res
}
