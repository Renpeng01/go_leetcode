package main

import "fmt"

// 原地排序
// 稳定排序
// 时间复杂度
// 最坏O(n²)
// 最好 O(n)
//
// 平均(加权平均期望时间)
// 有序度：数组中具有有序关系的元素对个数，对于一个完全有序的数组，它的有序度也叫满序度  满序度-逆序度=有序度
// 冒泡排序包含两个操作，比较和交换，每次交换，有序度加一，交换的次数是逆序度 n*(n-1)/2 - 初始有序度。
// 对于包含n个数据的数组进行冒泡排序，最坏情况下，初始状态的有序度是0，则要进行 n*(n-1)/2次交换，最好情况下，初始有序度是n*(n-1)/2。取中间值 n*(n-1)/2/2。则平均需要 n*(n-1)/2/2此交换操作
// 所以平均时间复杂度为O(n²)
func BubbleSort(data []int) []int {

	for i := len(data) - 1; i > 0; i-- {
		isSorted := true
		for j := i - 1; j >= 0; j-- {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return data
}

// 核心思想：动态的往有序集合中添加数据
// 将数组分为两个区间，已排序区间和未排序区间，初始已排序区间只有一个元素，取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序
// 原地排序 稳定排序
// 时间复杂度
func InsertSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0; j-- {
			if data[i] < data[j] {
				data[j], data[i] = data[i], data[j] // TODO 这里可以优化 见InsertSortV1
			}
		}
	}
	return data
}

func InsertSortV1(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 1; i < len(data); i++ {
		val := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if val < data[j] {
				data[j+1] = data[j] // 与InsertSort相比，少了一次赋值
			} else {
				break
			}
		}
		data[j+1] = val
	}
	return data
}

func main() {
	data := []int{1, 3, 5, 3, 8, 9, 5, 6, 7}
	res := InsertSortV1(data)
	fmt.Println("sorted: ", res)
}
