package main

// 时间复杂度 O(logn)
// 二分查找依赖顺序表结构，二分查找算法需要按照下表随机访问元素
// 二分查找针对的树有序数据
// 如果数据集合有频繁的插入和删除操作，想用二分查找，要么每次插入删除之后保证数据有序，要么在第二次查找之前都先进行排序，针对这种动态数据集合，无论哪种方法，维护有序的成本都很高
// 所以二分查找只能用在插入，删除操作不频繁，一次排序多次查找的场景中
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
