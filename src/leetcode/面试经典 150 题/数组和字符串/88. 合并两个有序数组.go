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

// 从后往前遍历，将最大的放在后面，可以节省一个slice的临时变量
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1Idx := m - 1
	nums2Idx := n - 1

	tail := m + n - 1
	for nums1Idx >= 0 && nums2Idx >= 0 {
		if nums1[nums1Idx] > nums2[nums2Idx] {
			nums1[tail] = nums1[nums1Idx]
			nums1Idx--
		} else {
			nums1[tail] = nums2[nums2Idx]
			nums2Idx--
		}
		tail--
	}

	for k := nums1Idx; k >= 0; k-- {
		nums1[tail] = nums1[k]
		tail--

	}
	for k := nums2Idx; k >= 0; k-- {
		nums1[tail] = nums2[k]
		tail--
	}
}
