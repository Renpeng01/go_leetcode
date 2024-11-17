package main

// 时间复杂度 O(logn)
// 二分查找依赖顺序表结构，二分查找算法需要按照下表随机访问元素
// 二分查找针对的树有序数据
// 如果数据集合有频繁的插入和删除操作，想用二分查找，要么每次插入删除之后保证数据有序，要么在第二次查找之前都先进行排序，针对这种动态数据集合，无论哪种方法，维护有序的成本都很高
// 所以二分查找只能用在插入，删除操作不频繁，一次排序多次查找的场景中
// 数据量太小不适合二分查找。不过有个例外，如果数据之间的比较操作非常耗时，不管数据量大小，都推荐使用二分查找，比如数组中存储的都是长度超过300的字符串
// 数据量太大也不适合二分查找，二分查找底层依赖数组，要求空间连续，对内存的要求比较苛刻
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

// ***************二分查找变形

// 1. 查找第一个值等于给定值的元素
func bsearch1(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)>>1

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if (a[mid-1] != value) || mid == 0 {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1

}

// 2. 查找最后一个值等于给定的元素
func bsearch2(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)>>1

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == len(a)-1 || a[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}

		}
	}
	return -1
}

// 3. 查找第一个大于等于给定值的元素
func bsearch3(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] >= value {
			if mid == 0 && a[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 4. 查找最后一个小于等于给定值的元素
func bsearch4(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] <= value {
			if mid == len(a)-1 || a[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}

	}
	return -1
}
