package main

// 堆是一种特殊的树
// 堆是一个完全二叉树 （完全二叉树比较适合用数组来存储，用数组来存储完全二叉树是非常节省存储空间的，因为不需要存储左右子节点的指针）
// 堆中每一个节点的值都必须大于等于（小于等于）其子树中的每个节点

type Heap struct {
	Vals  []int
	n     int
	count int
}

func (heap *Heap) Insert(data int) {
	if heap.count >= heap.n {
		return // 堆满
	}

	heap.count++
	heap.Vals[heap.count] = data
	i := heap.count
	for i/2 > 0 && heap.Vals[i] > heap.Vals[i/2] {
		swap(heap.Vals, i, i/2)
		i = i / 2
	}
}

func (heap *Heap) RemoveMax() {
	if heap.count == 0 {
		return
	}
	heap.Vals[1] = heap.Vals[heap.count]
	heap.heapify(heap.Vals, heap.count, 1)
}

func (heap *Heap) heapify(a []int, n, i int) {
	for {
		maxPos := i
		if i*2 < n && a[i] < a[i*2] {
			maxPos = i * 2
		}
		if i*2+1 < n && a[i] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
