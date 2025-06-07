package main

import "gthub.com/Renpeng01/go_leetcode/src/utils"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy1 := make([]int, len(nums1))
	for i, v := range nums1 {
		copy1[i] = v
	}

	nums1Idx := 0
	nums2Idx := 0
	i := 0
	for nums1Idx < m && nums2Idx < n {
		if copy1[nums1Idx] < nums2[nums2Idx] {
			nums1[i] = copy1[nums1Idx]
			nums1Idx++
		} else {
			nums1[i] = nums2[nums2Idx]
			nums2Idx++
		}
		i++
	}

	for k := nums1Idx; k < m; k++ {
		nums1[i] = copy1[k]
		i++

	}
	for k := nums2Idx; k < n; k++ {
		nums1[i] = nums2[k]
		i++
	}

	utils.PrintIntSlice(nums1)
}

func main() {
	num1 := []int{4, 5, 6, 0, 0, 0}
	num2 := []int{1, 2, 3}
	merge(num1, 3, num2, 3)
}
