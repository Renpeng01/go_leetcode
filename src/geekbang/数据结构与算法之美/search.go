package main

// 时间复杂度 O(logn)
// 二分查找依赖顺序表结构，二分查找算法需要按照下表随机访问元素
// 二分查找针对的树有序数据
func bsearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)>>1
		if value == a[mid] {
			return mid
		}
		if a[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
