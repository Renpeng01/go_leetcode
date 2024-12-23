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
// 最坏时间复杂度 O(n²)
// 最好时间复杂度 O(n)
// 平均时间复杂度 O(n²) 向数组插入的平均复杂度为O(n),则对于插入排序来说，每次插入操作相当于数组中插入一个数据，循环执行n次，故平均时间复杂度为O(n²)
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
				data[j+1] = data[j] // 优化点：与InsertSort相比，少了一次赋值
			} else {
				break // 优化点：如果有序，可以提前break
			}
		}
		data[j+1] = val
	}
	return data
}

// 原地排序，不稳定排序
// 时间复杂度 最好，最坏，平均都是O(n²)
func SelectSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}

// 冒泡排序和插入排序的时间复杂度都是O(n)²,都是原地稳定排序，为什么插入排序比冒泡更受欢迎
// 冒泡排序不管怎么优化，元素交换的次数是一个固定值，是原始数据的逆序度，插入排序也是，从代码实现上看，冒泡排序的赋值操作多余插入排序

// 稳定排序，非原地排序
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n) 在merge后，临时开辟的空间会释放
func MergeSort(data []int) {
	mSort(&data, 0, len(data)-1)
	return
}

func mSort(data *[]int, start, end int) {
	if end <= start {
		return
	}
	mid := start + (end-start)/2
	mSort(data, start, mid)
	mSort(data, mid+1, end)
	merge(data, (*data)[start:mid+1], (*data)[mid+1:end+1], start, end)
}

func merge(source *[]int, data1, data2 []int, start, end int) {
	i := 0
	j := 0
	res := make([]int, 0, len(data1)+len(data2))
	for i < len(data1) && j < len(data2) {
		if data1[i] < data2[j] {
			res = append(res, data1[i])
			i++
		} else {
			res = append(res, data2[j])
			j++
		}
	}

	for i < len(data1) {
		res = append(res, data1[i])
		i++
	}
	for j < len(data2) {
		res = append(res, data2[j])
		j++
	}

	for i, v := range res {
		(*source)[i+start] = v
	}
}

// 快速排序  原地 不稳定排序
// 原理：如果要排序数组中下标从p到r之间的一组数据，选择p到r之间任意一个数据作为pivot（分区点）
// 遍历p到r之间的数据，将小于pivot的放在左边，大于pivot的放在右边，将pivot放在中间。经过这一步之后，数组p到r之间的数据就被分成了3份，前面p到q-1之间都是小于pivot的，中间是pivot，后面q+1到r是大于pivot的
// 时间复杂度  如果每次分区，都正好把数组分成大小接近的两个小分区，则时间复杂度为O(nlogn)。在大部分情况下是O(nlogn)，在极端情况先是O(n²)
func QuickSort(data []int) {
	qSort(data, 0, len(data)-1)
}

func qSort(data []int, p, r int) {
	if p >= r {
		return
	}
	q := partation(data, 0, r)
	qSort(data, p, q-1)
	qSort(data, q+1, r)
}

func partation(data []int, p, r int) int {
	pivot := data[r]
	i := p
	for j := p; j <= r-1; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[r] = data[r], data[i]
	return i
}

// ************** 线性排序

// 桶排序
// 将要排序的数据分到几个有序桶，每个桶里的数据再单独进行排序，桶内排序完之后，再把每个桶里的数据按照顺序依次取出，组成的序列就是有序的
// 如果要排序的数据有n个，均匀的划分到m个桶内，每个桶里就有k=n/m个元素，每个桶内使用快排，时间复杂度是O(k*logk)。m个桶的排序时间复杂度就是 O(mk*logk)，因为k=n/m 所以O(n*log(n/m)),当桶的个数接近数据个数n时，log(n/m)就是一个非常小的常量，这时桶排序的时间复杂度接近O(n)
// 要求 1。 要排序的数据要很容易划分成m个桶，2. 桶与桶之间有天然的大小顺序。这样，每个桶内的数据都排完序后，桶与桶之间的数据不需要再进行排序。其次，各个桶之间的数据分布比较均匀。如果分布不均匀，那桶内数据排序的时间复杂度就不是常量级了

// 桶排序适合用在外部排序（数据量比较大，内存有限，无法将数据全部加载到内存 ）eg 10GB的订单数据，希望按照金额进行排序，但内存有限只有几百M，
// 先扫描一遍文件，看订单的范围，假设经过扫描之后得到金额最小1元，最大10万元，我们将所有订单根据金额划分到100个桶。第一个桶存1~1000元，第二个存1001~2000元依次类推，理想情况下，如果订单1到10万元均匀分布，订单会被均匀划分到100个文件中，每个文件100MB。将这些文件依次依次加载大内存排序
// 等所有文件都排完后，按照序号依次读取斌写入一个大文件中。
// 当1~10万不是均匀的，某个区间特别大，就需要继续划分
func BucketSort()

// 计数排序
// 计数排序是桶排序的一种特殊情况。
// 当要排序的n个数据范围不大，比如最大值k，就把数据分成k个桶，每个桶的数据值时相同的，省掉了桶内排序
// 只能用在数据范围不大的场景中

// 基数排序
// 假设有10万个手机号吗，希望将这10万个手机号从小打到排序。
// 比较手机号的a和b的大小，如果前面几位中，a的号已经大于b了，那后面的几位就用不着了，借助稳定排序算法，先按照最后一位来排序手机号码，然后倒数第二位，以此类推，经过11次排序后，手机号码就有序了。根据每一位来排序，可以用桶排序或技术排序

func main() {
	// sorted 1,3,3,5,5,6,7,8,9
	data := []int{1, 3, 5, 3, 8, 9, 5, 6, 7}
	QuickSort(data)
	fmt.Println("sorted: ", data)
}
