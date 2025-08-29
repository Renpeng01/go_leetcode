package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { //  这里是小于还是小于等于，需要模拟总结下(区间（左闭右闭 左闭右开）定义岁影响因素 )
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1
}
