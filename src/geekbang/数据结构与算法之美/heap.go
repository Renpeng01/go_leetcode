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

// 堆排序
// 1. 建堆 第一种，按照传统方式一个一个插入建立堆。从前往后处理数组数据，并且每个数据插入堆中时，都是从下往上堆化
//        第二种是从后往前处理数组，并且每个数据都是从上往下堆化。因为叶子节点往下堆化只能自己和自己比较，所以直接从最后一个非叶子节点开始堆化
// 2. 排序 数组中的第一个元素就是堆顶，也就是最大元素，把他和最后一个元素交换，然后再进行堆化

func buildHeap(a []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapify(a, n, i)
	}

}

// 每个节点堆化的时间复杂度是O(logn)，则n/2 + 1 个节点堆化总时间O(nlogn)，但是实际上是O(n)，因为叶子节点不需要堆化，所以需要堆化的节点从倒数第二层开始，每个节点堆化的过程中，需要比较和交换的节点个数，跟这个节点的高度k成正比
func heapify(a []int, n, i int) {
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

func heapSort(a []int, n int) {
	buildHeap(a, n)
	k := n
	for k > 1 {
		swap(a, 1, k)
		k--
		heapify(a, k, 1)
	}
}
