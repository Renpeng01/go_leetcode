package main

func findKthLargest(nums []int, k int) int {

	if len(nums) < k {
		return -1
	}
	nums = append([]int{0}, nums...)

	for i := 1; i < k; i++ {
		nums[i] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		buildMaxHeap(nums)
	}
	return nums[1]
}

func buildMaxHeap(nums []int) {
	n := (len(nums) - 1) / 2

	for i := n; i >= 0; i-- {
		// build(nums, i)
		if i*2 >= len(nums) {
			break
		}

		if i*2+1 >= len(nums) {
			if nums[i*2] > nums[i] {
				nums[i], nums[i*2] = nums[i*2], nums[i]
			}
			continue
		}
	}
}

func build(nums []int, index int) {
	if index*2 >= len(nums) {
		return
	}

	if index*2+1 >= len(nums) {
		if nums[index*2] > nums[index] {
			nums[index], nums[index*2] = nums[index*2], nums[index]
			return
		}
		return
	}

	if nums[index] > nums[index*2] && nums[index] > nums[index*2+1] {
		return
	}

	swipIndex := index * 2
	if nums[index*2+1] > nums[index*2] {
		swipIndex++
	}
	nums[index], nums[swipIndex] = nums[swipIndex], nums[index]
	build(nums, swipIndex)
}
